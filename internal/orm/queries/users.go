package queries

import (
	"github.com/gofrs/uuid"
	"github.com/google/martian/v3/log"
	"github.com/myminicommission/api/internal/orm"
	"github.com/myminicommission/api/internal/orm/models"
)

// GetUser retrieves a User record for the given ID
func GetUser(orm *orm.ORM, id uuid.UUID) (*models.User, error) {
	var user models.User
	db := orm.DB.New()
	db = db.First(&user, "id = ?", id.String())

	if db.Error != nil {
		log.Errorf("[ORM][users] %s", db.Error.Error())
		return nil, db.Error
	}

	return &user, nil
}

// CreateUser creates a User record
func CreateUser(orm *orm.ORM, nickname, name, email string) (*models.User, error) {
	user := models.User{
		NickName: &nickname,
		Name:     &name,
		Email:    email,
	}

	db := orm.DB.New()
	db = db.Create(&user)

	return &user, db.Error
}
