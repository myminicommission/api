package queries

import (
	"errors"

	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
	log "github.com/myminicommission/api/internal/logger"
	"github.com/myminicommission/api/internal/orm"
	"github.com/myminicommission/api/internal/orm/models"
)

// CreateGame creates a game with the given name and returns the new record
func CreateGame(orm *orm.ORM, name string) (*models.Game, error) {
	var game models.Game

	db := orm.DB.FirstOrCreate(&game, models.Game{Name: name})

	if db.Error != nil {
		return nil, db.Error
	}

	log.Info(game)

	return &game, nil
}

// GetGames returns an alpha sorted collection of Games
func GetGames(orm *orm.ORM) ([]*models.Game, error) {
	games := []*models.Game{}
	db := orm.DB.Order("name asc")
	db = db.Preload("Minis")
	db = db.Find(&games)

	if db.Error != nil {
		log.Errorf("[ORM][GetGames] %s", db.Error.Error())
		return nil, db.Error
	}

	return games, nil
}

// GetGame returns the first Game record with matching ID
func GetGame(orm *orm.ORM, id uuid.UUID) (*models.Game, error) {
	var game models.Game
	db := orm.DB.Preload("Minis")
	db = db.First(&game, "id = ?", id.String())

	if db.Error != nil {
		log.Errorf("[ORM][GetGame] %s", db.Error.Error())
		return nil, db.Error
	}

	return &game, nil
}

// GetGameMinis returns all GameMini records associated with the given Game ID
func GetGameMinis(orm *orm.ORM, id uuid.UUID) ([]*models.GameMini, error) {
	whereID := "game_id = ?"
	minis := []*models.GameMini{}
	db := orm.DB.Where(whereID, id.String())
	db = db.Find(&minis)

	if db.Error != nil {
		notFound := errors.Is(db.Error, gorm.ErrRecordNotFound)
		if notFound {
			log.Warnf("[ORM][GetGameMinis] %s", db.Error.Error())
			return nil, nil
		} else {
			log.Errorf("[ORM][GetGameMinis] %s", db.Error.Error())
			return nil, db.Error
		}
	}

	return minis, nil
}

// GetGameMini returns the first GameMini record with matching ID
func GetGameMini(orm *orm.ORM, id uuid.UUID) (*models.GameMini, error) {
	var mini models.GameMini
	db := orm.DB.First(&mini, "id = ?", id.String())

	if db.Error != nil {
		notFound := errors.Is(db.Error, gorm.ErrRecordNotFound)
		if notFound {
			log.Warnf("[ORM][GetGameMini] %s", db.Error.Error())
			return nil, nil
		} else {
			log.Errorf("[ORM][GetGameMini] %s", db.Error.Error())
			return nil, db.Error
		}
	}

	return &mini, nil
}

// GetMiniByNameAndGameName returns the first GameMini record with matching name and game
func GetMiniByNameAndGameName(orm *orm.ORM, name, gameName string) (*models.GameMini, error) {
	var mini models.GameMini

	db := orm.DB.Preload("Game")
	db = db.Joins("left join games on games.id = game_minis.game_id")
	db = db.Where("game_minis.name = ? AND games.name = ?", name, gameName)
	err := db.First(&mini).Error

	if err != nil {
		notFound := errors.Is(db.Error, gorm.ErrRecordNotFound)
		if notFound || err.Error() == gorm.ErrRecordNotFound.Error() {
			log.Warnf("[ORM][GetMiniByNameAndGameName] %s", err)
			return nil, gorm.ErrRecordNotFound
		} else {
			log.Errorf("[ORM][GetMiniByNameAndGameName] %s", err)
			return nil, db.Error
		}
	}

	return &mini, nil
}

// CreateGameMini creates a mini with the given name and game and returns the new record
func CreateGameMini(orm *orm.ORM, game uuid.UUID, name string) (*models.GameMini, error) {
	var mini models.GameMini

	db := orm.DB.FirstOrCreate(&mini, models.GameMini{
		Name:   name,
		GameID: game,
	})

	if db.Error != nil {
		notFound := errors.Is(db.Error, gorm.ErrRecordNotFound)
		if notFound {
			log.Warnf("[ORM][CreateGameMini] %s", db.Error.Error())
			return nil, nil
		} else {
			log.Errorf("[ORM][CreateGameMini] %s", db.Error.Error())
			return nil, db.Error
		}
	}

	log.Info(mini)

	return &mini, nil
}
