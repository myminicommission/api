package helpers

import (
	"errors"
	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
	"github.com/myminicommission/api/graph/helpers/transformations"
	"github.com/myminicommission/api/graph/model"
	"github.com/myminicommission/api/internal/orm"
	"github.com/myminicommission/api/internal/orm/models"
	"github.com/myminicommission/api/internal/orm/mutations"
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

// SaveMiniConfig saves a supplied MiniConfig for the given user.
// This will overwrite the existing record for the user
func SaveMiniConfig(orm *orm.ORM, input *model.MiniConfigInput, user *models.User) (*model.MiniConfig, error) {
	gameMini, err := queries.GetGameMini(orm, uuid.FromStringOrNil(input.Mini))
	if err != nil {
		return nil, err
	}
	dbConfig, err := queries.GetMiniConfig(orm, gameMini, user)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	// looks like this is a new record
	if dbConfig == nil || (err != nil && errors.Is(err, gorm.ErrRecordNotFound)) {
		dbConfig = &models.MiniConfig{
			Name: gameMini.Name,
			User: user,
			UserID: user.ID,
			Mini: gameMini,
			MiniID: gameMini.ID,
		}
	}

	dbConfig.Price = *input.Price
	dbConfig.Size = *input.Size

	// save the record
	dbConfig, err = mutations.SaveMiniConfig(orm, dbConfig)
	if err != nil {
		return nil, err
	}

	return transformations.DBMiniConfigToGQLMiniConfig(dbConfig)
}
