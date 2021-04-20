package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/mux"

	"github.com/myminicommission/api/graph"
	"github.com/myminicommission/api/graph/generated"
	log "github.com/myminicommission/api/internal/logger"
	"github.com/myminicommission/api/internal/middlewares/auth"
	"github.com/myminicommission/api/internal/orm"
	"github.com/myminicommission/api/internal/utils"
	"github.com/sirupsen/logrus"
	"github.com/urfave/negroni"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	var serverconf = &utils.ServerConfig{
		Port: utils.MustGet("PORT"),
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

	// negroni middleware
	n := negroni.Classic()

	// mux
	r := mux.NewRouter()

	// use some middlewares

	// auth middleware
	authMiddleware := auth.CreateMiddleware(serverconf, orm)
	r.Use(authMiddleware.Authorize)

	config := generated.Config{Resolvers: &graph.Resolver{
		ORM: orm,
	}}

	gqlSrv := handler.NewDefaultServer(generated.NewExecutableSchema(config))

	if serverconf.GraphQL.IsPlaygroundEnabled {
		r.Handle(serverconf.GraphQL.PlaygroundPath, playground.Handler("GraphQL playground", serverconf.GraphQL.Path))
		log.Infof("GraphQL playground enabled on http://localhost:%s%s", serverconf.Port, serverconf.GraphQL.PlaygroundPath)
	}
	r.Handle(serverconf.GraphQL.Path, gqlSrv)

	log.Infof("connect to http://localhost:%s%s for queries", serverconf.Port, serverconf.GraphQL.Path)

	n.UseHandler(r)
	srv := &http.Server{
		Handler:      n,
		Addr:         ":" + serverconf.Port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Error(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	_ = srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Info("shutting down")
	os.Exit(0)
}
