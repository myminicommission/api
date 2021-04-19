package models

import (
	"github.com/gofrs/uuid"
	"github.com/myminicommission/api/graph/model"
)

type MiniConfig struct {
	BaseModelSoftDelete
	Name   string
	Size   model.MiniSize
	Price  float64
	User   *User
	UserID uuid.UUID
	Mini   *GameMini
	MiniID uuid.UUID
}
