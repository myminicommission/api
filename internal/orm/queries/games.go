package queries

import (
	"github.com/gofrs/uuid"
	log "github.com/myminicommission/api/internal/logger"
	"github.com/myminicommission/api/internal/orm"
	"github.com/myminicommission/api/internal/orm/models"
)

// GetGames returns an alpha sorted collection of Games
func GetGames(orm *orm.ORM) ([]*models.Game, error) {
	games := []*models.Game{}
	db := orm.DB.New()
	db = db.Order("name asc")
	db = db.Preload("Minis")
	db = db.Find(&games)

	if db.Error != nil {
		log.Errorf("[ORM][games] %s", db.Error.Error())
		return nil, db.Error
	}

	return games, nil
}

// GetGame returns the first Game record with matching ID
func GetGame(orm *orm.ORM, id uuid.UUID) (*models.Game, error) {
	var game models.Game
	db := orm.DB.New()
	db = db.Preload("Minis")
	db = db.First(&game, "id = ?", id.String())

	if db.Error != nil {
		log.Errorf("[ORM][games] %s", db.Error.Error())
		return nil, db.Error
	}

	return &game, nil
}

// GetGameMinis returns all GameMini records associated with the given Game ID
func GetGameMinis(orm *orm.ORM, id uuid.UUID) ([]*models.GameMini, error) {
	whereID := "game_id = ?"
	minis := []*models.GameMini{}
	db := orm.DB.New()
	db = db.Where(whereID, id.String())
	db = db.Find(&minis)

	if db.Error != nil {
		log.Errorf("[ORM][game_minis] %s", db.Error.Error())
		return nil, db.Error
	}

	return minis, nil
}

// GetGameMini returns the first GameMini record with matching ID
func GetGameMini(orm *orm.ORM, id uuid.UUID) (*models.GameMini, error) {
	var mini models.GameMini
	db := orm.DB.New()
	db = db.First(&mini, "id = ?", id.String())

	if db.Error != nil {
		log.Errorf("[ORM][game_minis] %s", db.Error.Error())
		return nil, db.Error
	}

	return &mini, nil
}
