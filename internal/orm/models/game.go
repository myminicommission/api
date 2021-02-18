package models

import (
	"github.com/gofrs/uuid"
	"github.com/myminicommission/api/graph/model"
)

// Game has many GameMinis
type Game struct {
	BaseModelSoftDelete
	Name  string
	Minis []*GameMini
}

// GameMini belongs to Game, GameID is the foreign key
type GameMini struct {
	BaseModelSoftDelete
	Name   string
	Size   model.MiniSize
	Game   *Game
	GameID uuid.UUID
}
