package helpers

import (
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/myminicommission/api/graph/helpers/transformations"
	"github.com/myminicommission/api/graph/model"
	"github.com/myminicommission/api/internal/orm"
	"github.com/myminicommission/api/internal/orm/queries"
)

// GetGames returns a list of all games
func GetGames(orm *orm.ORM) ([]*model.Game, error) {
	dbRecords, err := queries.GetGames(orm)
	if err != nil {
		return nil, err
	}

	games := []*model.Game{}
	// transform the games
	for _, dbGame := range dbRecords {
		// transform the game for GQL
		game, err := transformations.DBGameToGQLGame(dbGame)
		if err != nil {
			return nil, err
		}

		games = append(games, game)
	}

	return games, nil
}

// GetGame returns a single game by ID
func GetGame(orm *orm.ORM, gameID uuid.UUID) (*model.Game, error) {
	dbGame, err := queries.GetGame(orm, gameID)
	if err != nil {
		return nil, err
	}

	// transform the game for GQL
	game, err := transformations.DBGameToGQLGame(dbGame)
	if err != nil {
		return nil, err
	}

	return game, nil
}

// GetGameMinis returns a single game by ID
func GetGameMinis(orm *orm.ORM, gameID uuid.UUID) ([]*model.GameMini, error) {
	dbRecords, err := queries.GetGameMinis(orm, gameID)
	if err != nil {
		return nil, err
	}

	if len(dbRecords) == 0 {
		return nil, fmt.Errorf("could not find minis for game with id: %s", gameID.String())
	}

	minis := []*model.GameMini{}
	for _, record := range dbRecords {
		mini, err := transformations.DBGameMiniToGQLGameMini(record)
		if err != nil {
			return nil, err
		}
		minis = append(minis, mini)
	}

	return minis, nil
}

// GetGameMini returns a single GameMini by ID
func GetGameMini(orm *orm.ORM, id uuid.UUID) (*model.GameMini, error) {
	dbGameMini, err := queries.GetGameMini(orm, id)
	if err != nil {
		return nil, err
	}

	mini, err := transformations.DBGameMiniToGQLGameMini(dbGameMini)
	if err != nil {
		return nil, err
	}

	return mini, nil
}
