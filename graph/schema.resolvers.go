package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/myminicommission/api/graph/generated"
	"github.com/myminicommission/api/graph/model"
)

func (r *queryResolver) Commissions(ctx context.Context, status *model.Status, artist *string, patron *string) ([]*model.Commission, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Commission(ctx context.Context, id string) (*model.Commission, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) SavedMinis(ctx context.Context, user string) ([]*model.SavedMini, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
