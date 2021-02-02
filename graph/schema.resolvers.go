package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/myminicommission/api/graph/generated"
	"github.com/myminicommission/api/graph/model"
	"github.com/myminicommission/api/internal/api/resolver"
)

func (r *mutationResolver) CreateEstimate(ctx context.Context, input model.NewEstimate) (*model.Estimate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateMini(ctx context.Context, input model.NewMini) (*model.Mini, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) SaveDefaultMiniCost(ctx context.Context, input model.NewDefaultMiniCost) (*model.DefaultMiniCost, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateQuote(ctx context.Context, input model.NewQuote) (*model.Quote, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateQuoteComment(ctx context.Context, input model.NewComment, quote string) (*model.QuoteComment, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateQuoteMiniQuantityComment(ctx context.Context, input model.NewComment, quote string, miniQuantity string) (*model.QuoteMiniQuantityComment, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Esitmate(ctx context.Context, input string) (*model.Estimate, error) {
	return resolver.GetEstimate(input)
}

func (r *queryResolver) Estimates(ctx context.Context) ([]*model.Estimate, error) {
	return resolver.GetEstimates()
}

func (r *queryResolver) EstimatesForUserID(ctx context.Context, input string) ([]*model.Estimate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Minis(ctx context.Context) ([]*model.Mini, error) {
	return resolver.GetMinis()
}

func (r *queryResolver) MinisWithName(ctx context.Context, input string) ([]*model.Mini, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Quote(ctx context.Context, input string) (*model.Quote, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Quotes(ctx context.Context) ([]*model.Quote, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) QuotesForUserID(ctx context.Context, input string) ([]*model.Quote, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
