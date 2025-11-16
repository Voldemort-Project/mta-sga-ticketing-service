package pgservice

import (
	"context"
	"errors"

	"github.com/Voldemort-Project/sga-service/src/domain/entities"
	"github.com/Voldemort-Project/sga-service/src/domain/repositories"
	pgmodels "github.com/Voldemort-Project/sga-service/src/infra/db/postgres/models"
	"gorm.io/gorm"
)

type roleService struct {
	db *gorm.DB
}

func NewRoleService(db *gorm.DB) repositories.RoleRepositoryImpl {
	return &roleService{db: db}
}

func (s *roleService) Upsert(ctx context.Context, data any, tx *gorm.DB) error {
	trx := s.db
	if tx != nil {
		trx = tx
	}

	payload, ok := data.(*entities.RoleEntity)
	if !ok {
		return errors.New("data is not a role entity")
	}

	model := pgmodels.RoleModel{}
	model.FromEntity(payload)

	if err := trx.
		WithContext(ctx).
		Session(&gorm.Session{FullSaveAssociations: true}).
		Save(&model).Error; err != nil {
		return err
	}
	return nil
}

func (s *roleService) Detail(ctx context.Context, id string, tx *gorm.DB) (any, error) {
	trx := s.db
	if tx != nil {
		trx = tx
	}

	model := pgmodels.RoleModel{}
	if err := trx.
		WithContext(ctx).
		Session(&gorm.Session{QueryFields: true}).
		Where("id = ?", id).
		First(&model).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return model.ToEntity(), nil
}

func (s *roleService) GetRoleByName(ctx context.Context, name string, tx *gorm.DB) (*entities.RoleEntity, error) {
	trx := s.db
	if tx != nil {
		trx = tx
	}

	model := pgmodels.RoleModel{}
	if err := trx.
		WithContext(ctx).
		Session(&gorm.Session{QueryFields: true}).
		Where("name = ?", name).
		First(&model).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return model.ToEntity(), nil
}
