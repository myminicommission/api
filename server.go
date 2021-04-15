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
	levelStr := utils.MustGet("LOG_LEVEL")
	level, err := logrus.ParseLevel(levelStr)
	// if there was an error parsing the level the default to Info
	if err != nil {
		log.Warnf("Could not parse [%s]. Defaulting to Info log level.", levelStr)
		level = logrus.InfoLevel
	}

	var serverconf = &utils.ServerConfig{
		Log: utils.LogConfig{
			Level: level,
		},
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
			LogMode:     utils.MustGetBool("GORM_LOGMODE"),
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

	logrus.SetLevel(serverconf.Log.Level)

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
	}
	http.Handle(serverconf.GraphQL.Path, srv)

	log.Infof("connect to http://localhost:%s/ for GraphQL playground", serverconf.Port)
	log.Fatal(http.ListenAndServe(":"+serverconf.Port, nil))
}
