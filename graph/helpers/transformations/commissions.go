package transformations

import (
	"github.com/myminicommission/api/graph/model"
	"github.com/myminicommission/api/internal/orm/models"
)

// DBCommissionToGQLCommission transforms a DB Commission to a GQL Commission
func DBCommissionToGQLCommission(i *models.Commission) (o *model.Commission, err error) {
	artist, err := DBUserToGQLUser(i.Artist)
	if err != nil {
		return
	}

	patron, err := DBUserToGQLUser(i.Patron)
	if err != nil {
		return
	}

	o = &model.Commission{
		ID:        i.ID.String(),
		Status:    i.Status,
		Total:     i.Total,
		Artist:    artist,
		Patron:    patron,
		CreatedAt: i.CreatedAt,
		UpdatedAt: i.UpdatedAt,
	}

	// transform the minis
	for _, dbMini := range i.Minis {
		mini, err := DBCommissionedMiniToGQLCommissionedMini(dbMini)
		if err != nil {
			return nil, err
		}
		// append the mini to the commission
		o.Minis = append(o.Minis, mini)
	}

	return
}

// DBCommissionedMiniToGQLCommissionedMini transforms a commission mini from db to gql
func DBCommissionedMiniToGQLCommissionedMini(i *models.CommissionedMini) (o *model.CommissionedMini, err error) {
	notes := i.Notes
	o = &model.CommissionedMini{
		ID:        i.ID.String(),
		Price:     i.Price,
		Size:      i.Size,
		Quantity:  i.Quantity,
		Notes:     &notes,
		CreatedAt: i.CreatedAt,
		UpdatedAt: i.UpdatedAt,
	}

	return
}
