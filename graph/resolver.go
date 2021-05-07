package graph

import (
	"context"

	"github.com/myminicommission/api/internal/middlewares/auth"
	"github.com/myminicommission/api/internal/orm"
	"github.com/myminicommission/api/internal/orm/models"
)

//go:generate go run github.com/99designs/gqlgen

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver holds the commons items for resolvers (think DI)
type Resolver struct {
	ORM *orm.ORM
}

func (r *Resolver) GetCurrentUser(ctx context.Context) *models.User {
	return ctx.Value(auth.UserCtxKey).(*models.User)
}
