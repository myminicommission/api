package jobs

import (
	"github.com/jinzhu/gorm"
	"github.com/myminicommission/api/graph/model"
	"github.com/myminicommission/api/internal/orm/models"
	"gopkg.in/gormigrate.v1"
)

var firstCommission *models.Commission = &models.Commission{
	Total:  50,
	Status: model.StatusQuote,
	Minis: []*models.CommissionedMini{
		{
			Name:     "Darth Vader",
			Size:     model.MiniSizeMedium,
			Price:    50,
			Quantity: 1,
			Notes:    "Operative pose",
		},
	},
}

// SeedCommissions populates the database with some starter data for commissions
var SeedCommissions *gormigrate.Migration = &gormigrate.Migration{
	ID: "SEED_COMMISSIONS",
	Migrate: func(db *gorm.DB) error {
		var artist models.User
		var patron models.User

		// get the artist and patron from the database
		err := db.First(&artist, "nick_name = ?", "TestUser1").Error
		if err != nil {
			return err
		}

		err = db.First(&patron, "nick_name = ?", "TestUser2").Error
		if err != nil {
			return err
		}

		// associate the commission with the users
		firstCommission.Artist = &artist
		firstCommission.Patron = &patron

		return db.Create(&firstCommission).Error
	},
	Rollback: func(db *gorm.DB) error {
		return db.Delete(&firstCommission).Error
	},
}
