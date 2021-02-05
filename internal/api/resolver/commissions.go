package resolver

import (
	"fmt"
	"time"

	"github.com/myminicommission/api/graph/model"
)

func getCommissions() []*model.Commission {
	return []*model.Commission{
		{
			ID:        "70bfc0e1-7927-49da-83e8-be8157b4a00e",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Total:     100,
			Minis: []*model.CommissionedMini{
				{
					ID:       "f96f9fe1-c6fa-463a-9f24-2a86666642c5",
					Quantity: 1,
					Name:     "Silent King",
					Size:     model.MiniSizeLarge,
					Price:    100,
				},
			},
			Artist: &model.User{
				Name:      "Extreme Moderation",
				ID:        "25d816e3-31bd-4cdf-86ab-e2bdb406b907",
				CreatedAt: time.Now().Add(-5 * time.Hour),
				UpdatedAt: time.Now().Add(-5 * time.Hour),
			},
		},
	}
}

// GetCommission returns a single record for the given ID
func GetCommission(id string) (*model.Commission, error) {
	var commission *model.Commission

	for _, comm := range getCommissions() {
		if comm.ID == id {
			commission = comm
		}
	}

	if commission == nil {
		return nil, fmt.Errorf("could not find commission with id %s", id)
	}

	return commission, nil
}

// GetCommissions returns commission objects for the given request
func GetCommissions() ([]*model.Commission, error) {
	return getCommissions(), nil
}
