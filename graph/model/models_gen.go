// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type DefaultMiniCost struct {
	ID   string   `json:"id"`
	Size MiniSize `json:"size"`
	Cost float64  `json:"cost"`
}

type DefaultMiniOptionCost struct {
	ID   string   `json:"id"`
	Size MiniSize `json:"size"`
	Name string   `json:"name"`
	Cost float64  `json:"cost"`
}

// Estimate type is automatically generated based on default stored values
type Estimate struct {
	ID        string          `json:"id"`
	CreatedAt time.Time       `json:"createdAt"`
	Minis     []*MiniQuantity `json:"minis"`
	Total     float64         `json:"total"`
	User      *User           `json:"user"`
}

type Mini struct {
	ID   string   `json:"id"`
	Name string   `json:"name"`
	Size MiniSize `json:"size"`
	Cost float64  `json:"cost"`
}

// Optional treatment for minis
type MiniOption struct {
	ID   string  `json:"id"`
	Name string  `json:"name"`
	Cost float64 `json:"cost"`
}

type MiniQuantity struct {
	ID       string        `json:"id"`
	Quantity int           `json:"quantity"`
	Mini     *Mini         `json:"mini"`
	Options  []*MiniOption `json:"options"`
}

type NewComment struct {
	Body string `json:"body"`
}

type NewDefaultMiniCost struct {
	Size MiniSize `json:"size"`
	Cost float64  `json:"cost"`
}

type NewEstimate struct {
	UserID string `json:"userId"`
}

type NewMini struct {
	Name string   `json:"name"`
	Size MiniSize `json:"size"`
}

type NewQuote struct {
	Estimate string `json:"estimate"`
}

type Quote struct {
	ID        string          `json:"id"`
	CreatedAt time.Time       `json:"createdAt"`
	Estimate  *Estimate       `json:"estimate"`
	Minis     []*MiniQuantity `json:"minis"`
	User      *User           `json:"user"`
}

type QuoteComment struct {
	ID        string    `json:"id"`
	Quote     *Quote    `json:"quote"`
	CreatedAt time.Time `json:"createdAt"`
	Body      string    `json:"body"`
	User      *User     `json:"user"`
}

type QuoteMiniQuantityComment struct {
	ID           string        `json:"id"`
	Quote        *Quote        `json:"quote"`
	MiniQuantity *MiniQuantity `json:"miniQuantity"`
	CreatedAt    time.Time     `json:"createdAt"`
	Body         string        `json:"body"`
	User         *User         `json:"user"`
}

type User struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
}

// Enum MiniSize
type MiniSize string

const (
	MiniSizeTiny    MiniSize = "TINY"
	MiniSizeRegular MiniSize = "REGULAR"
	MiniSizeLarge   MiniSize = "LARGE"
	MiniSizeTitanic MiniSize = "TITANIC"
)

var AllMiniSize = []MiniSize{
	MiniSizeTiny,
	MiniSizeRegular,
	MiniSizeLarge,
	MiniSizeTitanic,
}

func (e MiniSize) IsValid() bool {
	switch e {
	case MiniSizeTiny, MiniSizeRegular, MiniSizeLarge, MiniSizeTitanic:
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
