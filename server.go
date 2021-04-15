package main

import (
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/myminicommission/api/graph"
	"github.com/myminicommission/api/graph/generated"
	log "github.com/myminicommission/api/internal/logger"
	"github.com/myminicommission/api/internal/orm"
	"github.com/sirupsen/logrus"

	_ "github.com/joho/godotenv/autoload"
)

const defaultPort = "3001"

func main() {
	logrus.SetLevel(logrus.DebugLevel)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// ORM stuff
	orm, err := orm.Factory()
	if err != nil {
		log.Panic(err)
	}

	config := generated.Config{Resolvers: &graph.Resolver{
		ORM: orm,
	}}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(config))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Infof("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
