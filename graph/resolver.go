package graph

import (
	"github.com/myminicommission/api/internal/orm"
	"github.com/myminicommission/api/internal/orm/models"
)

//go:generate go run github.com/99designs/gqlgen

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	ORM *orm.ORM
}

func (r *Resolver) GetUser(nickName string) (*models.User, error) {
	var user models.User
	db := r.ORM.DB.New()
	db = db.First(&user, "nick_name = ?", "TestUser1")
	return &user, db.Error
}
