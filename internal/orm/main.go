package orm

import (
	"github.com/jinzhu/gorm"
	log "github.com/myminicommission/api/internal/logger"
	"github.com/myminicommission/api/internal/orm/migrations"
	"github.com/myminicommission/api/internal/utils"

	//Imports the database dialect of choice
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var autoMigrate, logMode bool
var dsn, dialect string

// ORM struct to hold the GORM pointer to the DB
type ORM struct {
	DB *gorm.DB
}

func init() {
	dialect = utils.MustGet("GORM_DIALECT")
	dsn = utils.MustGet("GORM_CONNECTION_DSN")
	logMode = utils.MustGetBool("GORM_LOGMODE")
	autoMigrate = utils.MustGetBool("GORM_AUTOMIGRATE")
}

// Factory creates a db connection with the selected dialect and DSN
func Factory() (*ORM, error) {
	db, err := gorm.Open(dialect, dsn)
	if err != nil {
		log.Panicf("[ORM] err: %s", err.Error())
	}

	orm := &ORM{
		DB: db,
	}

	// log every sql command
	db.LogMode(logMode)

	// automigrate
	if autoMigrate {
		err = migrations.ServiceAutoMigration(orm.DB)
	}

	log.Info("[ORM] Database connection initialized")
	return orm, err
}
