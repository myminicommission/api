package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/myminicommission/api/graph"
	"github.com/myminicommission/api/graph/generated"
	"github.com/myminicommission/api/graph/model"
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
	config.Directives.HasRole = func(ctx context.Context, obj interface{}, next graphql.Resolver, role model.Role) (interface{}, error) {
		// stubbed out hasRole value.
		// This clearly needs to be replaced with
		// something that knows who the current user is.
		hasRole := true
		if !hasRole {
			return nil, fmt.Errorf("missing role %v", role)
		}
		return next(ctx)
	}

	config.Directives.IsOwner = func(ctx context.Context, obj interface{}, next graphql.Resolver, isOwner bool) (interface{}, error) {
		// stubbed out validatedIsOwner value.
		// This clearly needs to be replaced with
		// something that knows who the current user is.
		validatedIsOwner := true
		if !validatedIsOwner {
			return nil, fmt.Errorf("authenticated user is not the owner")
		}
		return next(ctx)
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(config))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
