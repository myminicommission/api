package models

import (
	"github.com/gofrs/uuid"
	"github.com/myminicommission/api/graph/model"
)

// Commission type represents commissions in the database
type Commission struct {
	BaseModelSoftDelete
	Artist                    *User
	ArtistID                  uuid.UUID
	Patron                    *User
	PatronID                  uuid.UUID
	Status                    model.Status
	Minis                     []*CommissionedMini
	Total                     float64
	CommissionDiscussionItems []*CommissionDiscussionItem
}

func (c *Commission) DiscussionItems() []*DiscussionItem {
	var discussionItems []*DiscussionItem

	for _, cdi := range c.CommissionDiscussionItems {
		discussionItems = append(discussionItems, cdi.DiscussionItem)
	}

	return discussionItems
}

// CommissionDiscussionItem represents a relationship between Commission and DiscussionItem
type CommissionDiscussionItem struct {
	BaseModelSoftDelete
	Commission       *Commission
	CommissionID     uuid.UUID
	DiscussionItem   *DiscussionItem
	DiscussionItemID uuid.UUID
}

// DiscussionItem type represents a fragment of a conversation
type DiscussionItem struct {
	BaseModelSoftDelete
	Author   *User
	AuthorID uuid.UUID
	Body     string
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
