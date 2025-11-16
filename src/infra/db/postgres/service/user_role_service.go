package pgservice

import (
	"context"
	"errors"

	"github.com/Voldemort-Project/sga-service/src/domain/entities"
	"github.com/Voldemort-Project/sga-service/src/domain/repositories"
	pgmodels "github.com/Voldemort-Project/sga-service/src/infra/db/postgres/models"
	"gorm.io/gorm"
)

type userRoleService struct {
	db *gorm.DB
}

func NewUserRoleService(db *gorm.DB) repositories.UserRoleRepositoryImpl {
	return &userRoleService{db: db}
}

func (s *userRoleService) Upsert(ctx context.Context, data any, tx *gorm.DB) error {
	trx := s.db
	if tx != nil {
		trx = tx
	}

	payload, ok := data.(*entities.UserRoleEntity)
	if !ok {
		return errors.New("data is not a user role entity")
	}

	model := pgmodels.UserRoleModel{}
	model.FromEntity(payload)

	if err := trx.
		WithContext(ctx).
		Session(&gorm.Session{FullSaveAssociations: true}).
		Save(&model).Error; err != nil {
		return err
	}
	return nil
}

func (s *userRoleService) Detail(ctx context.Context, id string, tx *gorm.DB) (any, error) {
	trx := s.db
	if tx != nil {
		trx = tx
	}

	model := pgmodels.UserRoleModel{}
	if err := trx.
		WithContext(ctx).
		Session(&gorm.Session{QueryFields: true}).
		Where("id = ?", id).
		First(&model).Error; err != nil {
		return nil, err
	}
	return model.ToEntity(), nil
}
