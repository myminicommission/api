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

func (r *mutationResolver) CreateGame(ctx context.Context, name string) (*model.Game, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateGame(ctx context.Context, input model.GameInput) (*model.Game, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateGameMini(ctx context.Context, input *model.GameMiniInput) (*model.GameMini, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateGameMini(ctx context.Context, id string, input model.GameMiniInput) (*model.GameMini, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) MyCommissions(ctx context.Context) ([]*model.Commission, error) {
	// TODO: determine current user or reject request
	return helpers.MyCommissions(r.ORM, uuid.FromStringOrNil("af8f7135-07e5-47c8-bced-36b782b8ff64"))
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

func (r *queryResolver) Games(ctx context.Context) ([]*model.Game, error) {
	return helpers.GetGames(r.ORM)
}

func (r *queryResolver) Game(ctx context.Context, id string) (*model.Game, error) {
	return helpers.GetGame(r.ORM, uuid.FromStringOrNil(id))
}

func (r *queryResolver) GameMinis(ctx context.Context, game string) ([]*model.GameMini, error) {
	return helpers.GetGameMinis(r.ORM, uuid.FromStringOrNil(game))
}

func (r *queryResolver) GameMini(ctx context.Context, id string) (*model.GameMini, error) {
	return helpers.GetGameMini(r.ORM, uuid.FromStringOrNil(id))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
