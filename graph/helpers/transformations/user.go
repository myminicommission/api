package transformations

import (
	"github.com/myminicommission/api/graph/model"
	"github.com/myminicommission/api/internal/orm/models"
)

// DBUserToGQLUser transforms [user] db input to gql type
func DBUserToGQLUser(i *models.User) (o *model.User, err error) {
	o = &model.User{
		ID:        i.ID.String(),
		Name:      *i.Name,
		Nickname:  *i.NickName,
		CreatedAt: i.CreatedAt,
		UpdatedAt: i.UpdatedAt,
		Roles:     []*model.Role{},
	}

	if i.Picture != nil {
		o.Picture = *i.Picture
	}

	return
}
