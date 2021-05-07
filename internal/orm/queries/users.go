package queries

import (
	"github.com/gofrs/uuid"
	log "github.com/myminicommission/api/internal/logger"
	"github.com/myminicommission/api/internal/orm"
	"github.com/myminicommission/api/internal/orm/models"
)

// GetUser retrieves a User record for the given ID
func GetUser(orm *orm.ORM, id uuid.UUID) (*models.User, error) {
	var user models.User
	db := orm.DB.New()
	db = db.Preload("Socials")
	db = db.First(&user, "id = ?", id.String())

	if db.Error != nil {
		log.Errorf("[ORM][users] %s", db.Error.Error())
		return nil, db.Error
	}

	return &user, nil
}

// GetUserWithNickname returns the DB model for a User for the given nickname
func GetUserWithNickname(orm *orm.ORM, nickName string) (*models.User, error) {
	var user models.User
	db := orm.DB.New()
	db = db.Preload("Socials")
	db = db.First(&user, "LOWER(nick_name) = LOWER(?)", nickName)
	return &user, db.Error
}

// GetUserWithEmail returns the DB model for a User for the given email
func GetUserWithEmail(orm *orm.ORM, email string) (*models.User, error) {
	var user models.User
	db := orm.DB.New()
	db = db.Preload("socials")
	db = db.First(&user, "LOWER(email) = LOWER(?)", email)
	return &user, db.Error
}
