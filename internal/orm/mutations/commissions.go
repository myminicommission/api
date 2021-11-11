package mutations

import (
	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
	"github.com/myminicommission/api/graph/model"
	"github.com/myminicommission/api/internal/logger"
	"github.com/myminicommission/api/internal/orm"
	"github.com/myminicommission/api/internal/orm/models"
	"github.com/myminicommission/api/internal/orm/queries"
	"github.com/myminicommission/api/internal/utils"
)

var (
	log *logger.StandardLogger
)

func init() {
	log = logger.NewLogger()
}

// CreateCommission creates a commission with the supplied input
func CreateCommission(orm *orm.ORM, i *model.NewCommission, patronID uuid.UUID) (*models.Commission, error) {
	// get the artist
	artist, err := queries.GetUserWithNickname(orm, i.ArtistNickname)
	if err != nil {
		log.WithError(err).Errorf("error getting user with nickname %s", i.ArtistNickname)
		return nil, err
	}

	// get the patron
	patron, err := queries.GetUser(orm, patronID)
	if err != nil {
		log.WithError(err).Errorf("error getting user with id %s", patronID.String())
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
		Artist: artist,
		Patron: patron,
		Status: model.StatusQuote,
	}

	for _, miniInput := range i.Minis {
		log.Debug(miniInput.Name)
		// get the corresponding GameMini
		gameMini, err := queries.GetGameMini(orm, uuid.FromStringOrNil(miniInput.ID))
		if err != nil {
			log.WithError(err).Errorf("could not get mini with id %s (%s)", uuid.FromStringOrNil(miniInput.ID), miniInput.Name)
			return nil, err
		}
		log.Debug(gameMini.Name)

		// get the default price
		defaultPrice := utils.PriceForSize(&prices, miniInput.Size)

		// init a new CommissionedMini
		commissionedMini := models.CommissionedMini{
			Name:     gameMini.Name,
			Size:     miniInput.Size,
			Quantity: miniInput.Quantity,
			Price:    defaultPrice,
			Notes:    *miniInput.Notes,
		}
		log.Debug(commissionedMini.Name)

		// get the corresponding MiniConfig, if there is one
		miniConfig, err := queries.GetMiniConfig(orm, gameMini, artist)
		if err != nil {
			if isNotFound := err.Error() == gorm.ErrRecordNotFound.Error(); !isNotFound {
				log.WithError(err).Errorf("there was an error getting mini config for mini %s (%s)", gameMini.Name, gameMini.ID.String())
				return nil, err
			}

		}

		if miniConfig != nil && err == nil {
			// override commissionedMini values
			if miniConfig.Price >= 0 && miniConfig.Price != commissionedMini.Price {
				commissionedMini.Price = miniConfig.Price
			}

			if miniConfig.Size != commissionedMini.Size {
				commissionedMini.Size = miniConfig.Size
			}

			if miniConfig.Name != "" && miniConfig.Name != commissionedMini.Name {
				commissionedMini.Name = miniConfig.Name
			}
		}
		log.Debug(commissionedMini.Name)

		// append the mini to the commission
		commission.Minis = append(commission.Minis, &commissionedMini)

		// update the commission total
		commission.Total = commission.Total + float64(commissionedMini.Quantity)*commissionedMini.Price
	}

	// save the new commission
	err = orm.DB.Create(&commission).Error

	return &commission, err
}
