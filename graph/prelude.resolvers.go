package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql/introspection"
	"github.com/myminicommission/api/graph/generated"
)

func (r *DirectiveResolver) IsRepeatable(ctx context.Context, obj *introspection.Directive) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

// Directive returns generated.DirectiveResolver implementation.
func (r *Resolver) Directive() generated.DirectiveResolver { return &DirectiveResolver{r} }

type DirectiveResolver struct{ *Resolver }
