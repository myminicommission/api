package mutations

import (
	"github.com/gofrs/uuid"
	"github.com/myminicommission/api/graph/model"
	"github.com/myminicommission/api/internal/orm"
	"github.com/myminicommission/api/internal/orm/models"
	"github.com/myminicommission/api/internal/orm/queries"
	"github.com/myminicommission/api/internal/utils"
)

// CreateCommission creates a commission with the supplied input
func CreateCommission(orm *orm.ORM, i *model.NewCommission, patronID uuid.UUID) (*models.Commission, error) {
	// get the artist
	artist, err := queries.GetUser(orm, uuid.FromStringOrNil(i.Artist))
	if err != nil {
		return nil, err
	}

	// get the patron
	patron, err := queries.GetUser(orm, patronID)
	if err != nil {
		return nil, err
	}

	// TODO: get the artist's default pricing
	prices := model.Prices{
		Tiny:       9,
		Small:      9,
		Medium:     9,
		Large:      12,
		Extralarge: 15,
		Titanic:    30,
	}

	commission := models.Commission{
		Artist:              artist,
		Patron:              patron,
		Status:              model.StatusQuote,
	}

	for _, miniInput := range i.Minis {
		// get the corresponding GameMini
		gameMini, err := queries.GetGameMini(orm, uuid.FromStringOrNil(miniInput.ID))
		if err != nil {
			return nil, err
		}

		// get the default price
		defaultPrice := utils.PriceForSize(&prices, miniInput.Size)

		// init a new CommissionedMini
		commissionedMini := models.CommissionedMini{
			Name: gameMini.Name,
			Size: miniInput.Size,
			Quantity: miniInput.Quantity,
			Price: defaultPrice,
			Notes: *miniInput.Notes,
		}

		// get the corresponding MiniConfig, if there is one
		miniConfig, err := queries.GetMiniConfig(orm, gameMini, artist)
		if err != nil {
			return nil, err
		}


		if miniConfig != nil {
			// override commissionedMini values
			commissionedMini.Price = miniConfig.Price
			commissionedMini.Size = miniConfig.Size
			commissionedMini.Name = miniConfig.Name
		}

		// append the mini to the commission
		commission.Minis = append(commission.Minis, &commissionedMini)

		// update the commission total
		commission.Total = commission.Total + float64(commissionedMini.Quantity) * commissionedMini.Price
	}

	// save the new commission
	db := orm.DB.New()
	db = db.Create(&commission)

	return &commission, db.Error
}
