package migrations

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/myminicommission/api/internal/orm/migrations/jobs"
	"github.com/myminicommission/api/internal/orm/models"
	"gopkg.in/gormigrate.v1"
)

func updateMigration(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.User{},
		&models.Game{},
		&models.GameMini{},
		&models.Commission{},
		&models.CommissionedMini{},
	).Error
}

// ServiceAutoMigration migrates all the tables and modifications to the connected source
func ServiceAutoMigration(db *gorm.DB) error {
	// keep a lit of migrations
	m := gormigrate.New(db, gormigrate.DefaultOptions, nil)
	m.InitSchema(func(db *gorm.DB) error {
		log.Println("[Migration.InitSchema] Initializing database scheme")
		switch db.Dialect().GetName() {
		case "postgres":
			// lets create the UUID extension, the user has to have superuser permission
			db.Exec("create extension \"uuid-ossp\";")
		}

		if err := updateMigration(db); err != nil {
			return fmt.Errorf("[Migration.InitSchema]: %v", err)
		}
		// add more jobs below
		return nil
	})

	if err := m.Migrate(); err != nil {
		return err
	}

	if err := updateMigration(db); err != nil {
		return err
	}

	m = gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		jobs.SeedUsers,
		jobs.SeedGames,
		jobs.SeedCommissions,
	})

	return m.Migrate()
}
