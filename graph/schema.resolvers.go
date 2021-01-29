package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/myminicommission/api/graph/generated"
	"github.com/myminicommission/api/graph/model"
)

func (r *mutationResolver) CreateEstimate(ctx context.Context, input model.NewEstimate) (*model.Estimate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateMini(ctx context.Context, input model.NewMini) (*model.Mini, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Esitmate(ctx context.Context, input string) (*model.Estimate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Estimates(ctx context.Context) ([]*model.Estimate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) EstimatesForUserID(ctx context.Context, input string) ([]*model.Estimate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Minis(ctx context.Context) ([]*model.Mini, error) {
	silentKing := model.Mini{
		ID:   "c84b3f7a-b402-4bb4-a5eb-a08194e98aa3",
		Name: "Silent King",
		Size: model.MiniSizeLarge,
		Cost: 100,
	}
	minis := []*model.Mini{
		&silentKing,
	}
	return minis, nil
}

func (r *queryResolver) MinisWithName(ctx context.Context, input string) ([]*model.Mini, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
