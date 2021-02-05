// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

// Interface describing all minis
type Mini interface {
	IsMini()
}

// This is the heart of the application. Without this, nothing else matters.
type Commission struct {
	ID        string              `json:"id"`
	Artist    *User               `json:"artist"`
	Patron    *User               `json:"patron"`
	Status    Status              `json:"status"`
	Minis     []*CommissionedMini `json:"minis"`
	CreatedAt time.Time           `json:"createdAt"`
	UpdatedAt time.Time           `json:"updatedAt"`
	Total     float64             `json:"total"`
}

// Minis that belong to a commission. This is where the details of treatment are recorded as well as the price and quantity.
type CommissionedMini struct {
	ID        string    `json:"id"`
	Price     float64   `json:"price"`
	Quantity  int       `json:"quantity"`
	Notes     *string   `json:"notes"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Name      string    `json:"name"`
	Size      MiniSize  `json:"size"`
}

func (CommissionedMini) IsMini() {}

// Games are collections of minis
type Game struct {
	ID        string      `json:"id"`
	Name      string      `json:"name"`
	Minis     []*GameMini `json:"minis"`
	CreatedAt time.Time   `json:"createdAt"`
	UpdatedAt time.Time   `json:"updatedAt"`
}

// Represents a miniature. This type represents all miniatures across the system.
type GameMini struct {
	ID        string    `json:"id"`
	Game      *Game     `json:"game"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Name      string    `json:"name"`
	Size      MiniSize  `json:"size"`
}

func (GameMini) IsMini() {}

// Default prices as configured by the user
type Prices struct {
	ID         string  `json:"id"`
	TIny       float64 `json:"tINY"`
	SMall      float64 `json:"sMALL"`
	MEdium     float64 `json:"mEDIUM"`
	LArge      float64 `json:"lARGE"`
	EXtralarge float64 `json:"eXTRALARGE"`
	TItanic    float64 `json:"tITANIC"`
	User       *User   `json:"user"`
}

// Saved mini configuration. This is used to override the default pricing for a specific mini.
type SavedMini struct {
	ID        string    `json:"id"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	User      *User     `json:"user"`
	Name      string    `json:"name"`
	Size      MiniSize  `json:"size"`
}

func (SavedMini) IsMini() {}

// Basic user representation.
type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Roles     []*Role   `json:"roles"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Nickname  string    `json:"nickname"`
}

// Enum MiniSize
type MiniSize string

const (
	MiniSizeTiny       MiniSize = "TINY"
	MiniSizeSmall      MiniSize = "SMALL"
	MiniSizeMedium     MiniSize = "MEDIUM"
	MiniSizeLarge      MiniSize = "LARGE"
	MiniSizeExtralarge MiniSize = "EXTRALARGE"
	MiniSizeTitanic    MiniSize = "TITANIC"
)

var AllMiniSize = []MiniSize{
	MiniSizeTiny,
	MiniSizeSmall,
	MiniSizeMedium,
	MiniSizeLarge,
	MiniSizeExtralarge,
	MiniSizeTitanic,
}

func (e MiniSize) IsValid() bool {
	switch e {
	case MiniSizeTiny, MiniSizeSmall, MiniSizeMedium, MiniSizeLarge, MiniSizeExtralarge, MiniSizeTitanic:
		return true
	}
	return false
}

func (e MiniSize) String() string {
	return string(e)
}

func (e *MiniSize) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = MiniSize(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid MiniSize", str)
	}
	return nil
}

func (e MiniSize) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Enum Role
type Role string

const (
	RoleAdmin  Role = "ADMIN"
	RoleArtist Role = "ARTIST"
)

var AllRole = []Role{
	RoleAdmin,
	RoleArtist,
}

func (e Role) IsValid() bool {
	switch e {
	case RoleAdmin, RoleArtist:
		return true
	}
	return false
}

func (e Role) String() string {
	return string(e)
}

func (e *Role) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Role(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Role", str)
	}
	return nil
}

func (e Role) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Enum Status
type Status string

const (
	StatusEstimate   Status = "ESTIMATE"
	StatusQuote      Status = "QUOTE"
	StatusAccepted   Status = "ACCEPTED"
	StatusWaiting    Status = "WAITING"
	StatusInProgress Status = "IN_PROGRESS"
	StatusShipped    Status = "SHIPPED"
	StatusComplete   Status = "COMPLETE"
)

var AllStatus = []Status{
	StatusEstimate,
	StatusQuote,
	StatusAccepted,
	StatusWaiting,
	StatusInProgress,
	StatusShipped,
	StatusComplete,
}

func (e Status) IsValid() bool {
	switch e {
	case StatusEstimate, StatusQuote, StatusAccepted, StatusWaiting, StatusInProgress, StatusShipped, StatusComplete:
		return true
	}
	return false
}

func (e Status) String() string {
	return string(e)
}

func (e *Status) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Status(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Status", str)
	}
	return nil
}

func (e Status) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
