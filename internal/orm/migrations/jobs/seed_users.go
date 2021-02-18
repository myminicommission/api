package jobs

import (
	"github.com/jinzhu/gorm"
	"github.com/myminicommission/api/internal/orm/models"
	"gopkg.in/gormigrate.v1"
)

var (
	uname1                 = "Test User 1"
	email1                 = "testuser1@localhost"
	nname1                 = "TestUser1"
	firstUser *models.User = &models.User{
		Email:    email1,
		Name:     &uname1,
		NickName: &nname1,
	}

	uname2                  = "Test User 2"
	email2                  = "testuser2@localhost"
	nname2                  = "TestUser2"
	secondUser *models.User = &models.User{
		Email:    email2,
		Name:     &uname2,
		NickName: &nname2,
	}
)

// SeedUsers populates the database with some starting data
var SeedUsers *gormigrate.Migration = &gormigrate.Migration{
	ID: "SEED_USERS",
	Migrate: func(db *gorm.DB) error {
		err := db.Create(&firstUser).Error
		if err != nil {
			return err
		}

		err = db.Create(&secondUser).Error
		if err != nil {
			return err
		}
		return nil
	},
	Rollback: func(db *gorm.DB) error {
		err := db.Delete(&firstUser).Error
		if err != nil {
			return err
		}

		err = db.Delete(&secondUser).Error
		if err != nil {
			return err
		}
		return nil
	},
}
