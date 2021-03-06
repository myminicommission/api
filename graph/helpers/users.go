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
