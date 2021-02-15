package helpers

import (
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/google/martian/v3/log"
	"github.com/myminicommission/api/graph/helpers/transformations"
	"github.com/myminicommission/api/graph/model"
	"github.com/myminicommission/api/internal/orm"
	"github.com/myminicommission/api/internal/orm/models"
)

// GetUser returns a single user by ID
func GetUser(orm *orm.ORM, userID uuid.UUID) (*model.User, error) {
	entity := "users"
	whereID := "id = ?"
	dbRecords := []*models.User{}
	db := orm.DB.New()
	db = db.Where(whereID, userID.String())
	db = db.Find(&dbRecords)

	if db.Error != nil {
		log.Errorf("[ORM][%s] %s", entity, db.Error.Error())
		return nil, db.Error
	}

	if len(dbRecords) == 0 {
		return nil, fmt.Errorf("could not find user with id: %s", userID.String())
	}

	return transformations.DBUserToGQLUser(dbRecords[0])
}
