package models

// User defines a user for the app
type User struct {
	BaseModelSoftDelete
	Email    string  `gorm:"not null; unique_index:idx_email"`
	UserID   *string // external user ID
	Name     *string
	NickName *string
}
