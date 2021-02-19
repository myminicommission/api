package transformations

import (
	"github.com/myminicommission/api/graph/model"
	"github.com/myminicommission/api/internal/orm/models"
)

// DBMiniConfigToGQLMiniConfig transforms the MiniConfig from the DB model to the GQL model
func DBMiniConfigToGQLMiniConfig(i *models.MiniConfig) (o *model.MiniConfig, err error) {
	user, err := DBUserToGQLUser(i.User)
	if err != nil {
		return
	}

	mini, err := DBGameMiniToGQLGameMini(i.Mini)
	if err != nil {
		return
	}

	o = &model.MiniConfig{
		ID:        i.ID.String(),
		UpdatedAt: i.UpdatedAt,
		CreatedAt: i.CreatedAt,
		Name:      i.Name,
		Price:     i.Price,
		Size:      i.Size,
		User:      user,
		Mini:      mini,
	}
	return
}
