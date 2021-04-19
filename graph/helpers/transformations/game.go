package transformations

import (
	"github.com/myminicommission/api/graph/model"
	"github.com/myminicommission/api/internal/orm/models"
)

// DBGameToGQLGame transforms [game] db input to gql type
func DBGameToGQLGame(i *models.Game) (o *model.Game, err error) {
	o = &model.Game{
		ID:        i.ID.String(),
		CreatedAt: i.CreatedAt,
		UpdatedAt: i.UpdatedAt,
		Name:      i.Name,
	}

	for _, record := range i.Minis {
		mini, err := DBGameMiniToGQLGameMini(record)
		if err != nil {
			return nil, err
		}
		o.Minis = append(o.Minis, mini)
	}

	return
}

// DBGameMiniToGQLGameMini transforms [GameMini] db input tp gql type
func DBGameMiniToGQLGameMini(i *models.GameMini) (o *model.GameMini, err error) {
	o = &model.GameMini{
		ID:        i.ID.String(),
		CreatedAt: i.CreatedAt,
		UpdatedAt: i.UpdatedAt,
		Name:      i.Name,
		Size:      i.Size,
	}

	return
}
