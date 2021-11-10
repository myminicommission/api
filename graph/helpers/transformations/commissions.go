package transformations

import (
	"github.com/myminicommission/api/graph/model"
	"github.com/myminicommission/api/internal/logger"
	"github.com/myminicommission/api/internal/orm/models"
)

// DBCommissionToGQLCommission transforms a DB Commission to a GQL Commission
func DBCommissionToGQLCommission(i *models.Commission) (o *model.Commission, err error) {
	logger.Debug(i)
	artist, err := DBUserToGQLUser(i.Artist)
	if err != nil {
		logger.Errorf("error marshalling artist: %v - %v", i.Artist, err)
		return
	}

	patron, err := DBUserToGQLUser(i.Patron)
	if err != nil {
		logger.Errorf("error marshalling patron: %v - %v", i.Patron, err)
		return
	}

	o = &model.Commission{
		ID:        i.ID.String(),
		Status:    i.Status,
		Total:     i.Total,
		Artist:    artist,
		Patron:    patron,
		CreatedAt: i.CreatedAt,
		UpdatedAt: i.UpdatedAt,
	}

	println(o)

	// transform the discussion items
	for _, di := range i.DiscussionItems() {
		discussionItem, err := DBDiscussionItemToGQLDiscussionItem(di)
		if err != nil {
			return nil, err
		}
		// append the discussion item to the commission
		o.DiscussionItems = append(o.DiscussionItems, discussionItem)
	}

	// transform the minis
	for _, dbMini := range i.Minis {
		mini, err := DBCommissionedMiniToGQLCommissionedMini(dbMini)
		if err != nil {
			return nil, err
		}
		// append the mini to the commission
		o.Minis = append(o.Minis, mini)
	}

	return
}

// DBDiscussionItemToGQLDiscussionItem transforms a discussion item from the db to gql
func DBDiscussionItemToGQLDiscussionItem(i *models.DiscussionItem) (o *model.DiscussionItem, err error) {
	author, err := DBUserToGQLUser(i.Author)
	if err != nil {
		logger.Errorf("error marshalling author: %v - %v", i.Author, err)
		return
	}

	o = &model.DiscussionItem{
		ID:        i.ID.String(),
		CreatedAt: i.CreatedAt,
		UpdatedAt: i.UpdatedAt,
		Author:    author,
		Body:      i.Body,
	}

	return
}

// DBCommissionedMiniToGQLCommissionedMini transforms a commission mini from db to gql
func DBCommissionedMiniToGQLCommissionedMini(i *models.CommissionedMini) (o *model.CommissionedMini, err error) {
	notes := i.Notes
	o = &model.CommissionedMini{
		ID:        i.ID.String(),
		Name:      i.Name,
		Price:     i.Price,
		Size:      i.Size,
		Quantity:  i.Quantity,
		Notes:     &notes,
		CreatedAt: i.CreatedAt,
		UpdatedAt: i.UpdatedAt,
	}

	return
}
