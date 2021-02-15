package jobs

import (
	"github.com/jinzhu/gorm"
	"github.com/myminicommission/api/internal/orm/models"
	"gopkg.in/gormigrate.v1"
)

var (
	uname                  = "Test User"
	email                  = "test@user.com"
	nname                  = "TestUser123"
	firstUser *models.User = &models.User{
		Email:    email,
		Name:     &uname,
		NickName: &nname,
	}
)

var SeedUsers *gormigrate.Migration = &gormigrate.Migration{
	ID: "SEED_USERS",
	Migrate: func(db *gorm.DB) error {
		return db.Create(&firstUser).Error
	},
	Rollback: func(db *gorm.DB) error {
		return db.Delete(&firstUser).Error
	},
}
