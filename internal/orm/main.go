package orm

import (
	"context"

	log "github.com/myminicommission/api/internal/logger"
	"github.com/myminicommission/api/internal/orm/migrations"
	"github.com/myminicommission/api/internal/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ORM struct to hold the GORM pointer to the DB
type ORM struct {
	DB *gorm.DB
}

// Factory creates a db connection with the selected dialect and DSN
func Factory(config *utils.ServerConfig) (*ORM, error) {
	entry := config.Logger.Logger.WithContext(context.Background())
	db, err := gorm.Open(postgres.Open(config.Database.DSN), &gorm.Config{
		Logger: NewGormLogger(entry),
	})
	if err != nil {
		log.Panicf("[ORM] err: %s", err.Error())
	}

	orm := &ORM{
		DB: db,
	}

	// lets create the UUID extension, the user has to have superuser permission
	db.Exec("create extension \"uuid-ossp\";")

	// automigrate
	if config.Database.AutoMigrate {
		err = migrations.ServiceAutoMigration(orm.DB)
	}

	log.Info("[ORM] Database connection initialized")
	return orm, err
}
