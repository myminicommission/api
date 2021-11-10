package mutations

import (
	"github.com/myminicommission/api/internal/orm"
	"github.com/myminicommission/api/internal/orm/models"
	"github.com/myminicommission/api/internal/orm/queries"
)

// SaveMiniConfig creates or updates a MiniConfig record
func SaveMiniConfig(orm *orm.ORM, i *models.MiniConfig) (*models.MiniConfig, error) {
	db := orm.DB.Save(i)
	if db.Error != nil {
		return nil, db.Error
	}

	// get and return the saved MiniConfig
	return queries.GetMiniConfig(orm, i.Mini, i.User)
}
