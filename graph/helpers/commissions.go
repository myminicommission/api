package helpers

import (
	"github.com/gofrs/uuid"
	"github.com/myminicommission/api/graph/helpers/transformations"
	"github.com/myminicommission/api/graph/model"
	"github.com/myminicommission/api/internal/orm"
	"github.com/myminicommission/api/internal/orm/mutations"
	"github.com/myminicommission/api/internal/orm/queries"
)

// MyCommissions returns the commissions where
// the user is either the artist or the patron
func MyCommissions(orm *orm.ORM, id uuid.UUID) ([]*model.Commission, error) {
	// get the commissions
	dbCommissions, err := queries.GetMyCommissions(orm, id)
	if err != nil {
		return nil, err
	}

	var commissions []*model.Commission
	for _, dbCommission := range dbCommissions {
		// transform the dbCommission
		commission, err := transformations.DBCommissionToGQLCommission(dbCommission)
		if err != nil {
			return nil, err
		}

		// append the commission to the collection
		commissions = append(commissions, commission)
	}

	return commissions, nil
}

// GetCommission retrieves a commission from the DB by its ID
func GetCommission(orm *orm.ORM, id uuid.UUID) (*model.Commission, error) {
	dbCommission, err := queries.GetCommission(orm, id)
	if err != nil {
		return nil, err
	}

	return transformations.DBCommissionToGQLCommission(dbCommission)
}

// NewCommission creates a new commission from the input
func NewCommission(orm *orm.ORM, input *model.NewCommission, patronID uuid.UUID) (*model.Commission, error) {
	dbCommission, err := mutations.CreateCommission(orm, input, patronID)
	if err != nil {
		return nil, err
	}

	return transformations.DBCommissionToGQLCommission(dbCommission)
}