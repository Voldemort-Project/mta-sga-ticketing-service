package pgservice

import (
	"context"
	"errors"

	"github.com/Voldemort-Project/sga-service/src/domain/entities"
	"github.com/Voldemort-Project/sga-service/src/domain/repositories"
	pgmodels "github.com/Voldemort-Project/sga-service/src/infra/db/postgres/models"
	"gorm.io/gorm"
)

type roomService struct {
	db *gorm.DB
}

func NewRoomService(db *gorm.DB) repositories.RoomRepositoryImpl {
	return &roomService{db: db}
}

func (s *roomService) Upsert(ctx context.Context, data any, tx *gorm.DB) error {
	trx := s.db
	if tx != nil {
		trx = tx
	}

	payload, ok := data.(*entities.RoomEntity)
	if !ok {
		return errors.New("data is not a room entity")
	}

	model := pgmodels.RoomModel{}
	model.FromEntity(payload)

	if err := trx.
		WithContext(ctx).
		Session(&gorm.Session{FullSaveAssociations: true}).
		Save(&model).Error; err != nil {
		return err
	}
	return nil
}

func (s *roomService) Detail(ctx context.Context, id string, tx *gorm.DB) (any, error) {
	trx := s.db
	if tx != nil {
		trx = tx
	}

	model := pgmodels.RoomModel{}
	if err := trx.
		WithContext(ctx).
		Session(&gorm.Session{QueryFields: true}).
		Where("id = ?", id).
		First(&model).Error; err != nil {
		return nil, err
	}

	return model.ToEntity(), nil
}

func (s *roomService) GetRoomByNameAndAvailibility(ctx context.Context, name string, tx *gorm.DB) (*entities.RoomEntity, error) {
	trx := s.db
	if tx != nil {
		trx = tx
	}

	model := pgmodels.RoomModel{}
	if err := trx.
		WithContext(ctx).
		Session(&gorm.Session{QueryFields: true, PrepareStmt: true}).
		Where("name = ?", name).
		First(&model).Error; err != nil {
		return nil, err
	}

	return model.ToEntity(), nil
}

func (s *roomService) UpdateVisibility(ctx context.Context, id string, isAvailable bool, tx *gorm.DB) error {
	trx := s.db
	if tx != nil {
		trx = tx
	}
	if err := trx.WithContext(ctx).
		Model(&pgmodels.RoomModel{}).
		Where("id = ?", id).
		Update("is_available", isAvailable).Error; err != nil {
		return err
	}
	return nil
}
