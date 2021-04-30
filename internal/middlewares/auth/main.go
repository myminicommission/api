package auth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	log "github.com/myminicommission/api/internal/logger"
	"github.com/myminicommission/api/internal/orm"
	"github.com/myminicommission/api/internal/orm/models"
	"github.com/myminicommission/api/internal/orm/mutations"
	"github.com/myminicommission/api/internal/orm/queries"
	"github.com/myminicommission/api/internal/utils"

	"github.com/gofrs/uuid"
)

type Auth0OpenID struct {
	Issuer                            string   `json:"issuer"`
	AuthorizationEndpoint             string   `json:"authorization_endpoint"`
	TokenEndpoint                     string   `json:"token_endpoint"`
	UserinfoEndpoint                  string   `json:"userinfo_endpoint"`
	MfaChallengeEndpoint              string   `json:"mfa_challenge_endpoint"`
	JwksURI                           string   `json:"jwks_uri"`
	RegistrationEndpoint              string   `json:"registration_endpoint"`
	RevocationEndpoint                string   `json:"revocation_endpoint"`
	ScopesSupported                   []string `json:"scopes_supported"`
	ResponseTypesSupported            []string `json:"response_types_supported"`
	CodeChallengeMethodsSupported     []string `json:"code_challenge_methods_supported"`
	ResponseModesSupported            []string `json:"response_modes_supported"`
	SubjectTypesSupported             []string `json:"subject_types_supported"`
	IDTokenSigningAlgValuesSupported  []string `json:"id_token_signing_alg_values_supported"`
	TokenEndpointAuthMethodsSupported []string `json:"token_endpoint_auth_methods_supported"`
	ClaimsSupported                   []string `json:"claims_supported"`
	RequestURIParameterSupported      bool     `json:"request_uri_parameter_supported"`
	DeviceAuthorizationEndpoint       string   `json:"device_authorization_endpoint"`
}

type Auth0JWKS struct {
	Keys []struct {
		Alg string   `json:"alg"`
		Kty string   `json:"kty"`
		Use string   `json:"use"`
		N   string   `json:"n"`
		E   string   `json:"e"`
		Kid string   `json:"kid"`
		X5T string   `json:"x5t"`
		X5C []string `json:"x5c"`
	} `json:"keys"`
}

type AuthMiddleware struct {
	auth0Conf *utils.AuthProvider
	ORM       *orm.ORM
}

type ctxKey struct {
	name string
}

var userCtxKey = &ctxKey{name: "user"}

func CreateMiddleware(serverConfig *utils.ServerConfig, orm *orm.ORM) AuthMiddleware {
	m := AuthMiddleware{
		ORM: orm,
	}

	// extract the Auth0 configs
	for _, provider := range serverConfig.AuthProviders {
		if provider.Provider == "auth0" {
			m.auth0Conf = &provider
			continue
		}
	}

	if m.auth0Conf == nil {
		log.Fatal("could not find auth0 provider config")
	}

	return m
}

func (m *AuthMiddleware) Authorize(next http.Handler) http.Handler {
	log.Debug("[Middleware] AuthMiddleware.Authorize")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// setup a place to hold the user
		var user *models.User

		// extract the auth header
		authHeader := r.Header.Get("Authorization")
		authHeader = strings.Replace(authHeader, "Bearer ", "", 1)

		// if the token is present (ie, not blank), decode it
		if authHeader != "" {
			log.Debugf("Auth Token: %s", authHeader)

			var claims jwt.MapClaims

			// parse the token
			token, err := jwt.Parse(authHeader, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}

				// extract the claims
				claims = token.Claims.(jwt.MapClaims)

				// check the audience claims
				checkAud := claims.VerifyAudience(m.auth0Conf.ClientKey, false)
				if !checkAud {
					return nil, errors.New("invalid audience")
				}

				// check the issuer claims
				checkIss := claims.VerifyIssuer(fmt.Sprintf("https://%s/", m.auth0Conf.Domain), false)
				if !checkIss {
					return nil, errors.New("invalid issuer")
				}

				key, err := getAuth0JWKS(m.auth0Conf.Domain, token.Header["kid"].(string))
				if err != nil {
					return nil, err
				}

				cert := "-----BEGIN CERTIFICATE-----\n" + key + "\n-----END CERTIFICATE-----"
				result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))

				return result, nil
			})

			if err != nil {
				log.Errorf("Error parsing token: %s", err.Error())
			}

			if token.Valid {
				// if valid, lookup user in the DB and set them in context...
				email := fmt.Sprint(claims["email"])
				nickname := fmt.Sprint(claims["nickname"])
				name := fmt.Sprint(claims["name"])
				picture := fmt.Sprint(claims["picture"])

				// get the user (if there is one)
				user, err = queries.GetUserWithNickname(m.ORM, nickname)
				if err != nil {
					// handle the error
					log.Errorf("Error looking up user: %s", err.Error())

					// create the user if one wasn't found
					if err.Error() == "record not found" {
						user, err = mutations.CreateUser(m.ORM, &models.User{
							Name:     &name,
							NickName: &nickname,
							Email:    email,
							Picture:  &picture,
						})
						if err != nil {
							// handle the erorr
							log.Errorf("Error while creating the user: %s", err.Error())
						}
					}
				}
			}
		}

		if user.ID != uuid.Nil {
			ctx := context.WithValue(r.Context(), userCtxKey, user)
			r = r.WithContext(ctx)
		}

		// continue with the request
		next.ServeHTTP(w, r)
	})
}

func getAuth0JWKS(domain, kid string) (string, error) {
	openID, err := getAuth0OpenID(domain)
	if err != nil {
		return "", err
	}

	jwksURI := openID.JwksURI
	resp, err := http.Get(jwksURI)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var jwks Auth0JWKS
	err = json.Unmarshal(body, &jwks)
	if err != nil {
		return "", err
	}

	var x5c string
	for _, key := range jwks.Keys {
		if key.Kid == kid && len(key.X5C) == 1 {
			x5c = key.X5C[0]
		}
	}

	if x5c == "" {
		return "", fmt.Errorf("no key found for kid [%s]", kid)
	}

	return x5c, nil
}

func getAuth0OpenID(domain string) (Auth0OpenID, error) {
	var auth0OpenID Auth0OpenID

	openIDEndpoint := fmt.Sprintf("https://%s/.well-known/openid-configuration", domain)
	resp, err := http.Get(openIDEndpoint)
	if err != nil {
		log.Errorf("Error getting OpenID Configuration: %s", err.Error())
		return auth0OpenID, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return auth0OpenID, err
	}

	err = json.Unmarshal(body, &auth0OpenID)
	if err != nil {
		return auth0OpenID, err
	}

	return auth0OpenID, nil
}
