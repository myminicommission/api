package jobs

import (
	"github.com/jinzhu/gorm"
	"github.com/myminicommission/api/graph/model"
	"github.com/myminicommission/api/internal/orm/models"
	"gopkg.in/gormigrate.v1"
)

var firstGames []*models.Game = []*models.Game{
	{
		Name: "Star Wars Legion",
		Minis: []*models.GameMini{
			{
				Name: "Darth Vader",
				Size: model.MiniSizeMedium,
			},
		},
	},
	{
		Name: "Warhammer 40,000",
		Minis: []*models.GameMini{
			{
				Name: "Drukhari: Drazhar",
				Size: model.MiniSizeMedium,
			},
			{
				Name: "Drukhari: Raider",
				Size: model.MiniSizeLarge,
			},
		},
	},
}

// SeedGames sets up the inital games seed for the DB
var SeedGames *gormigrate.Migration = &gormigrate.Migration{
	ID: "SEED_GAMES",
	Migrate: func(db *gorm.DB) error {
		for _, game := range firstGames {
			err := db.Create(&game).Error
			if err != nil {
				return err
			}
		}
		return nil
	},
	Rollback: func(db *gorm.DB) error {
		for _, game := range firstGames {
			err := db.Delete(&game).Error
			if err != nil {
				return err
			}
		}
		return nil
	},
}
