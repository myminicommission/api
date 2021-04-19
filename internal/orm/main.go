package orm

import (
	"github.com/jinzhu/gorm"
	log "github.com/myminicommission/api/internal/logger"
	"github.com/myminicommission/api/internal/orm/migrations"
	"github.com/myminicommission/api/internal/utils"

	//Imports the database dialect of choice
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// ORM struct to hold the GORM pointer to the DB
type ORM struct {
	DB *gorm.DB
}

// Factory creates a db connection with the selected dialect and DSN
func Factory(config *utils.ServerConfig) (*ORM, error) {
	db, err := gorm.Open(config.Database.Dialect, config.Database.DSN)
	if err != nil {
		log.Panicf("[ORM] err: %s", err.Error())
	}

	orm := &ORM{
		DB: db,
	}

	// log every sql command if configured
	db.LogMode(config.Database.LogMode)

	// automigrate
	if config.Database.AutoMigrate {
		err = migrations.ServiceAutoMigration(orm.DB)
	}

	log.Info("[ORM] Database connection initialized")
	return orm, err
}
