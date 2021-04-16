package main

import (
	"net/http"
	"strings"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/myminicommission/api/graph"
	"github.com/myminicommission/api/graph/generated"
	log "github.com/myminicommission/api/internal/logger"
	"github.com/myminicommission/api/internal/orm"
	"github.com/myminicommission/api/internal/utils"
	"github.com/sirupsen/logrus"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	var serverconf = &utils.ServerConfig{
		Port:          utils.MustGet("SERVER_PORT"),
		SessionSecret: utils.MustGet("SESSION_SECRET"),
		JWT: utils.JWTConfig{
			Secret:    utils.MustGet("AUTH_JWT_SECRET"),
			Algorithm: utils.MustGet("AUTH_JWT_SIGNING_ALGORITHM"),
		},
		GraphQL: utils.GQLConfig{
			Path:                utils.MustGet("GQL_SERVER_GRAPHQL_PATH"),
			PlaygroundPath:      utils.MustGet("GQL_SERVER_GRAPHQL_PLAYGROUND_PATH"),
			IsPlaygroundEnabled: utils.MustGetBool("GQL_SERVER_GRAPHQL_PLAYGROUND_ENABLED"),
		},
		Database: utils.DBConfig{
			Dialect:     utils.MustGet("GORM_DIALECT"),
			DSN:         utils.MustGet("GORM_CONNECTION_DSN"),
			LogMode:     log.IsLevelEnabled(logrus.DebugLevel),
			AutoMigrate: utils.MustGetBool("GORM_AUTOMIGRATE"),
		},
		AuthProviders: []utils.AuthProvider{
			{
				Provider:  "auth0",
				ClientKey: utils.MustGet("PROVIDER_AUTH0_KEY"),
				Secret:    utils.MustGet("PROVIDER_AUTH0_SECRET"),
				Domain:    utils.MustGet("PROVIDER_AUTH0_DOMAIN"),
				Scopes:    strings.Split(utils.MustGet("PROVIDER_AUTH0_SCOPES"), ","),
			},
		},
	}

	// ORM stuff
	orm, err := orm.Factory(serverconf)
	if err != nil {
		log.Panic(err)
	}

	config := generated.Config{Resolvers: &graph.Resolver{
		ORM: orm,
	}}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(config))

	if serverconf.GraphQL.IsPlaygroundEnabled {
		http.Handle(serverconf.GraphQL.PlaygroundPath, playground.Handler("GraphQL playground", serverconf.GraphQL.Path))
		log.Infof("GraphQL playground enabled on http://localhost:%s%s", serverconf.Port, serverconf.GraphQL.PlaygroundPath)
	}
	http.Handle(serverconf.GraphQL.Path, srv)

	log.Infof("connect to http://localhost:%s%s for queries", serverconf.Port, serverconf.GraphQL.Path)
	log.Fatal(http.ListenAndServe(":"+serverconf.Port, nil))
}
