package queries

import (
	"github.com/gofrs/uuid"
	log "github.com/myminicommission/api/internal/logger"
	"github.com/myminicommission/api/internal/orm"
	"github.com/myminicommission/api/internal/orm/models"
	"gorm.io/gorm/clause"
)

// GetMyCommissions returns commissions where the
// supplied user is either the artist or the patron
func GetMyCommissions(orm *orm.ORM, id uuid.UUID) ([]*models.Commission, error) {
	var commissions []*models.Commission

	db := orm.DB.Where("artist_id = ?", id.String())
	db = db.Or("patron_id = ?", id.String())
	db = db.Preload(clause.Associations)
	db = db.Preload("CommissionDiscussionItems." + clause.Associations)
	db = db.Preload("CommissionDiscussionItems.DiscussionItem." + clause.Associations)
	db = db.Find(&commissions)

	if db.Error != nil {
		log.Errorf("[ORM][commissions] %s", db.Error.Error())
		return nil, db.Error
	}

	return commissions, nil
}

// GetCommission returns the commission with the provided ID
func GetCommission(orm *orm.ORM, id uuid.UUID) (*models.Commission, error) {
	var commission models.Commission

	db := orm.DB.Preload(clause.Associations)
	db = db.Preload("CommissionDiscussionItems." + clause.Associations)
	db = db.Preload("CommissionDiscussionItems.DiscussionItem." + clause.Associations)
	db = db.First(&commission, "id = ?", id.String())

	if db.Error != nil {
		log.Errorf("[ORM][commissions] %s", db.Error.Error())
		return nil, db.Error
	}

	return &commission, nil
}

// GetCommissionedMinis returns the minis associated with a commission
func GetCommissionedMinis(orm *orm.ORM, id uuid.UUID) ([]*models.CommissionedMini, error) {
	var minis []*models.CommissionedMini

	db := orm.DB.Where("commission_id = ?", id.String())
	db = db.Find(&minis)

	if db.Error != nil {
		log.Errorf("[ORM][commissioned_minis] %s", db.Error.Error())
		return nil, db.Error
	}

	return minis, nil
}

// GetCommissionedMini returns the commissioned mini with the given id
func GetCommissionedMini(orm *orm.ORM, id uuid.UUID) (*models.CommissionedMini, error) {
	var mini models.CommissionedMini

	db := orm.DB.First(&mini, "id = ?", id.String())

	if db.Error != nil {
		log.Errorf("[ORM][commissioned_minis] %s", db.Error.Error())
		return nil, db.Error
	}

	return &mini, nil
}
