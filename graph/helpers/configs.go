package helpers

import (
	"github.com/myminicommission/api/graph/helpers/transformations"
	"github.com/myminicommission/api/graph/model"
	"github.com/myminicommission/api/internal/orm"
	"github.com/myminicommission/api/internal/orm/models"
	"github.com/myminicommission/api/internal/orm/queries"
)

// GetMiniConfigs returns the MiniConfig objects for the given user
func GetMiniConfigs(orm *orm.ORM, user *models.User) ([]*model.MiniConfig, error) {
	dbConfigs, err := queries.GetMiniConfigs(orm, user)
	if err != nil {
		return nil, err
	}

	var configs []*model.MiniConfig
	for _, dbConfig := range dbConfigs {
		config, err := transformations.DBMiniConfigToGQLMiniConfig(dbConfig)
		if err != nil {
			return nil, err
		}
		configs = append(configs, config)
	}

	return configs, nil
}
