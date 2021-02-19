package jobs

import (
	"github.com/jinzhu/gorm"
	"github.com/myminicommission/api/graph/model"
	"github.com/myminicommission/api/internal/orm/models"
	"gopkg.in/gormigrate.v1"
)

var firstUserMiniConfig *models.MiniConfig = &models.MiniConfig{
	Name:  "Darth Vader",
	Size:  model.MiniSizeMedium,
	Price: 50,
}

var SeedMiniConfigs *gormigrate.Migration = &gormigrate.Migration{
	ID: "SEED_MINI_CONFIGS",
	Migrate: func(db *gorm.DB) error {
		var user models.User
		var mini models.GameMini

		// get the user from the DB
		err := db.First(&user, "nick_name = ?", "TestUser1").Error
		if err != nil {
			return err
		}

		// get the mini from the DB
		err = db.First(&mini, "name = ?", "Darth Vader").Error
		if err != nil {
			return err
		}

		firstUserMiniConfig.User = &user
		firstUserMiniConfig.Mini = &mini

		return db.Create(&firstUserMiniConfig).Error
	},
	Rollback: func(db *gorm.DB) error {
		return db.Delete(&firstUserMiniConfig).Error
	},
}
