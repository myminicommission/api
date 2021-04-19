package helpers

import (
	"github.com/gofrs/uuid"
	"github.com/myminicommission/api/graph/helpers/transformations"
	"github.com/myminicommission/api/graph/model"
	"github.com/myminicommission/api/internal/orm"
	"github.com/myminicommission/api/internal/orm/queries"
)

// GetUser returns a single user by ID
func GetUser(orm *orm.ORM, userID uuid.UUID) (*model.User, error) {
	user, err := queries.GetUser(orm, userID)
	if err != nil {
		return nil, err
	}

	return transformations.DBUserToGQLUser(user)
}

// CreateUser creates and returns a single User
func CreateUser(orm *orm.ORM, user *model.User) (*model.User, error) {
	println(user.Name)
	name := user.Name
	if name == "" {
		name = user.Nickname
	}

	newUser, err := queries.CreateUser(orm, user.Nickname, name)
	if err != nil {
		return nil, err
	}

	return transformations.DBUserToGQLUser(newUser)
}
