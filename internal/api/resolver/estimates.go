package resolver

import (
	"fmt"
	"time"

	"github.com/myminicommission/api/graph/model"
)

func getEstimates() []*model.Estimate {
	return []*model.Estimate{
		{
			ID:        "70bfc0e1-7927-49da-83e8-be8157b4a00e",
			CreatedAt: time.Now(),
			Total:     100,
			Minis: []*model.MiniQuantity{
				{
					ID:       "f96f9fe1-c6fa-463a-9f24-2a86666642c5",
					Quantity: 1,
					Mini: &model.Mini{
						ID:   "c84b3f7a-b402-4bb4-a5eb-a08194e98aa3",
						Name: "Silent King",
						Size: model.MiniSizeLarge,
						Cost: 100,
					},
					Options: []*model.MiniOption{
						{
							ID:   "dd3855d3-c73f-4a4d-abb4-24733a083ad8",
							Name: "Basing",
							Cost: 1,
						},
					},
				},
			},
			User: &model.User{
				Name:      "Extreme Moderation",
				ID:        "25d816e3-31bd-4cdf-86ab-e2bdb406b907",
				CreatedAt: time.Now().Add(-5 * time.Hour),
			},
		},
	}
}

// GetEstimate returns a single record for the given ID
func GetEstimate(id string) (*model.Estimate, error) {
	var estimate *model.Estimate

	for _, est := range getEstimates() {
		if est.ID == id {
			estimate = est
		}
	}

	if estimate == nil {
		return nil, fmt.Errorf("could not find estimate with id %s", id)
	}

	return estimate, nil
}

// GetEstimates returns estimate objects for the given request
func GetEstimates() ([]*model.Estimate, error) {
	return getEstimates(), nil
}
