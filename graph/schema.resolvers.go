package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/myminicommission/api/graph/generated"
	"github.com/myminicommission/api/graph/helpers"
	"github.com/myminicommission/api/graph/model"
)

func (r *mutationResolver) NewCommission(ctx context.Context, input model.NewCommission) (*model.Commission, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateCommission(ctx context.Context, input model.CommissionInput) (*model.Commission, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) SaveMiniConfig(ctx context.Context, input model.MiniConfigInput) (*model.MiniConfig, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Commissions(ctx context.Context, status *model.Status, artist *string, patron *string) ([]*model.Commission, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Commission(ctx context.Context, id string) (*model.Commission, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	return helpers.GetUser(r.ORM, uuid.FromStringOrNil(id))
}

func (r *queryResolver) MiniConfigs(ctx context.Context, user string) ([]*model.MiniConfig, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
