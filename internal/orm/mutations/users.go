package mutations

import (
	"github.com/myminicommission/api/graph/model"
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

// UpdateProfile changes the given user's ForHire and Name value to the input specified
func UpdateProfile(orm *orm.ORM, user *models.User, input *model.ProfileInput) error {
	user.ForHire = input.ForHire
	user.Name = &input.Name

	socials := user.Socials

	if input.Socials.Facebook != nil && *input.Socials.Facebook != "" {
		socials.Facebook = input.Socials.Facebook
	}

	if input.Socials.Instagram != nil && *input.Socials.Instagram != "" {
		socials.Instagram = input.Socials.Instagram
	}

	if input.Socials.Twitch != nil && *input.Socials.Twitch != "" {
		socials.Twitch = input.Socials.Twitch
	}

	if input.Socials.Twitter != nil && *input.Socials.Twitter != "" {
		socials.Twitter = input.Socials.Twitter
	}

	user.Socials = socials

	db := orm.DB.New()
	db = db.Save(user)

	return db.Error
}
