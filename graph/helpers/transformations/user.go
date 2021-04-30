package transformations

import (
	"github.com/myminicommission/api/graph/model"
	"github.com/myminicommission/api/internal/orm/models"
)

// DBUserToGQLUser transforms [user] db input to gql type
func DBUserToGQLUser(i *models.User) (*model.User, error) {
	o := model.User{
		ID:        i.ID.String(),
		Name:      *i.Name,
		Nickname:  *i.NickName,
		CreatedAt: i.CreatedAt,
		UpdatedAt: i.UpdatedAt,
		Roles:     []*model.Role{},
		Socials:   &model.Socials{},
		ForHire:   i.ForHire,
	}

	if i.Picture != nil {
		o.Picture = *i.Picture
	}

	if i.Socials != nil {
		o.Socials.Facebook = i.Socials.Facebook
		o.Socials.Instagram = i.Socials.Instagram
		o.Socials.Twitch = i.Socials.Twitch
		o.Socials.Twitter = i.Socials.Twitter
	}

	return &o, nil
}
