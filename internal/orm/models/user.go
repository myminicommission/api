package models

import "github.com/gofrs/uuid"

// User defines a user for the app
type User struct {
	BaseModelSoftDelete
	Email    string  `gorm:"not null; unique_index:idx_email"`
	UserID   *string // external user ID
	Name     *string
	NickName *string `gorm:"not null; unique_index:idx_nickname"`
	Picture  *string
	ForHire  bool `gorm:"not null; default:false"`
	Socials  *Socials
}

// Socials is a container to hold links to various social media stuff
type Socials struct {
	BaseModelSoftDelete
	User      User
	UserID    uuid.UUID
	Facebook  *string
	Instagram *string
	Twitch    *string
	Twitter   *string
}
