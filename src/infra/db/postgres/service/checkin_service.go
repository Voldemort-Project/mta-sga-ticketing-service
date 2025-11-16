package pgservice

import (
	"context"
	"errors"

	"github.com/Voldemort-Project/sga-service/src/domain/entities"
	"github.com/Voldemort-Project/sga-service/src/domain/repositories"
	pgmodels "github.com/Voldemort-Project/sga-service/src/infra/db/postgres/models"
	"gorm.io/gorm"
)

type checkinService struct {
	db *gorm.DB
}

func NewCheckinService(db *gorm.DB) repositories.CheckinRepositoryImpl {
	return &checkinService{db: db}
}

func (s *checkinService) Upsert(ctx context.Context, data any, tx *gorm.DB) error {
	trx := s.db
	if tx != nil {
		trx = tx
	}

	payload, ok := data.(*entities.CheckinEntity)
	if !ok {
		return errors.New("data is not a checkin entity")
	}
	model := pgmodels.CheckinModel{}
	model.FromEntity(payload)

	if err := trx.
		WithContext(ctx).
		Session(&gorm.Session{FullSaveAssociations: true}).
		Save(&model).Error; err != nil {
		return err
	}
	return nil
}

func (s *checkinService) Detail(ctx context.Context, id string, tx *gorm.DB) (any, error) {
	trx := s.db
	if tx != nil {
		trx = tx
	}

	model := pgmodels.CheckinModel{}
	if err := trx.
		WithContext(ctx).
		Session(&gorm.Session{QueryFields: true}).
		Where("id = ?", id).
		First(&model).Error; err != nil {
		return nil, err
	}

	return model.ToEntity(), nil
}
