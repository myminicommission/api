package models

import (
	"time"

	"github.com/gofrs/uuid"
)

// BaseModel defines the common columns that all db structs hold
// usually, db structs based on this have no soft delete
type BaseModel struct {
	// ID should use uuid_generate_v4() for the PK's
	ID        uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time `gorm:"index;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"index;not null;default:CURRENT_TIMESTAMP"`
}

// BaseModelSoftDelete defines the common columns that all db structs should hold.
// This struct also defines the fields for GORM triggers to
// detect if the entity should soft delete
type BaseModelSoftDelete struct {
	BaseModel
	DeletedAt *time.Time `sql:"index"`
}
