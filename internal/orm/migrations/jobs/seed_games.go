package jobs

import (
	"github.com/jinzhu/gorm"
	"github.com/myminicommission/api/graph/model"
	"github.com/myminicommission/api/internal/orm/models"
	"gopkg.in/gormigrate.v1"
)

var firstGame *models.Game = &models.Game{
	Name: "Star Wars Legion",
	Minis: []*models.GameMini{
		{
			Name: "Darth Vader",
			Size: model.MiniSizeMedium,
		},
	},
}

// SeedGames sets up the inital games seed for the DB
var SeedGames *gormigrate.Migration = &gormigrate.Migration{
	ID: "SEED_GAMES",
	Migrate: func(db *gorm.DB) error {
		return db.Create(&firstGame).Error
	},
	Rollback: func(db *gorm.DB) error {
		return db.Delete(&firstGame).Error
	},
}
