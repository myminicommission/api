package mutations

import (
	"github.com/myminicommission/api/internal/orm"
	"github.com/myminicommission/api/internal/orm/models"
)

// CreateUser creates a User record
func CreateUser(orm *orm.ORM, user *models.User) (*models.User, error) {
	if user.Socials == nil {
		user.Socials = &models.Socials{}
	}

	db := orm.DB.New()
	db = db.Create(user)

	return user, db.Error
}
