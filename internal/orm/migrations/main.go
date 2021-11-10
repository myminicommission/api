package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/myminicommission/api/internal/orm/models"
	"gorm.io/gorm"
)

// ServiceAutoMigration migrates all the tables and modifications to the connected source
func ServiceAutoMigration(db *gorm.DB) error {
	// keep a list of migrations
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		// users
		{
			ID: "users_table",
			Migrate: func(d *gorm.DB) error {
				return d.AutoMigrate(&models.User{})
			},
			Rollback: func(d *gorm.DB) error {
				return d.Migrator().DropTable("users")
			},
		},
		// games
		{
			ID: "games_table",
			Migrate: func(d *gorm.DB) error {
				return d.AutoMigrate(&models.Game{})
			},
			Rollback: func(d *gorm.DB) error {
				return d.Migrator().DropTable("games")
			},
		},
		// 		&models.GameMini{},
		{
			ID: "game_minis_table",
			Migrate: func(d *gorm.DB) error {
				return d.AutoMigrate(&models.GameMini{})
			},
			Rollback: func(d *gorm.DB) error {
				return d.Migrator().DropTable("game_minis")
			},
		},
		// 		&models.Commission{},
		{
			ID: "commissions_table",
			Migrate: func(d *gorm.DB) error {
				return d.AutoMigrate(&models.Commission{})
			},
			Rollback: func(d *gorm.DB) error {
				return d.Migrator().DropTable("commissions")
			},
		},
		// 		&models.DiscussionItem{},
		{
			ID: "discussion_items_table",
			Migrate: func(d *gorm.DB) error {
				return d.AutoMigrate(&models.DiscussionItem{})
			},
			Rollback: func(d *gorm.DB) error {
				return d.Migrator().DropTable("discussion_items")
			},
		},
		// 		&models.CommissionDiscussionItem{},
		{
			ID: "commission_discussion_items_table",
			Migrate: func(d *gorm.DB) error {
				return d.AutoMigrate(&models.CommissionDiscussionItem{})
			},
			Rollback: func(d *gorm.DB) error {
				return d.Migrator().DropTable("commission_discussion_items")
			},
		},
		// 		&models.CommissionedMini{},
		{
			ID: "commissioned_minis_table",
			Migrate: func(d *gorm.DB) error {
				return d.AutoMigrate(&models.CommissionedMini{})
			},
			Rollback: func(d *gorm.DB) error {
				return d.Migrator().DropTable("commissioned_minis")
			},
		},
		// 		&models.MiniConfig{},
		{
			ID: "mini_configs_table",
			Migrate: func(d *gorm.DB) error {
				return d.AutoMigrate(&models.MiniConfig{})
			},
			Rollback: func(d *gorm.DB) error {
				return d.Migrator().DropTable("mini_configs")
			},
		},
		// 		&models.Socials{},
		{
			ID: "socials_table",
			Migrate: func(d *gorm.DB) error {
				return d.AutoMigrate(&models.Socials{})
			},
			Rollback: func(d *gorm.DB) error {
				return d.Migrator().DropTable("socials")
			},
		},
	})

	return m.Migrate()
}
