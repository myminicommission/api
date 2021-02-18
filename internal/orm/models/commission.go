package models

import (
	"github.com/gofrs/uuid"
	"github.com/myminicommission/api/graph/model"
)

// Commission type represents commissions in the database
type Commission struct {
	BaseModelSoftDelete
	Artist   *User
	ArtistID uuid.UUID
	Patron   *User
	PatronID uuid.UUID
	Status   model.Status
	Minis    []*CommissionedMini
	Total    float64
}

// CommissionedMini belongs to a Commission
type CommissionedMini struct {
	BaseModelSoftDelete
	Name         string
	Size         model.MiniSize
	Price        float64
	Quantity     int
	Notes        string
	Commission   *Commission
	CommissionID uuid.UUID
}
