package queries

import (
	"github.com/myminicommission/api/internal/orm"
	"github.com/myminicommission/api/internal/orm/models"
)

// GetMiniConfigs gets MiniConfig records from the DB for the given user
func GetMiniConfigs(orm *orm.ORM, user *models.User) ([]*models.MiniConfig, error) {
	var configs []*models.MiniConfig

	db := orm.DB.New()
	db = db.Where("user_id = ?", user.ID.String())
	db = db.Preload("Mini")
	db = db.Preload("User")
	db = db.Find(&configs)

	return configs, db.Error
}
