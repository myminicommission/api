package graph

import "github.com/myminicommission/api/internal/orm"

//go:generate go run github.com/99designs/gqlgen

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	ORM *orm.ORM
}
