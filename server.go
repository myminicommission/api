package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/myminicommission/api/graph"
	"github.com/myminicommission/api/graph/generated"
	"github.com/myminicommission/api/internal/orm"

	_ "github.com/joho/godotenv/autoload"
)

const defaultPort = "8080"

func main() {
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

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
