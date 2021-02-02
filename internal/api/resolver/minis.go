package resolver

import "github.com/myminicommission/api/graph/model"

// GetMinis returns a list of minis
func GetMinis() ([]*model.Mini, error) {
	silentKing := model.Mini{
		ID:   "c84b3f7a-b402-4bb4-a5eb-a08194e98aa3",
		Name: "Silent King",
		Size: model.MiniSizeLarge,
		Cost: 100,
	}
	minis := []*model.Mini{
		&silentKing,
	}
	return minis, nil
}
