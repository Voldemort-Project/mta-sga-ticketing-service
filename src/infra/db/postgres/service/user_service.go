package pgservice

import (
	"context"
	"errors"

	"github.com/Voldemort-Project/sga-service/src/domain/entities"
	"github.com/Voldemort-Project/sga-service/src/domain/repositories"
	pgmodels "github.com/Voldemort-Project/sga-service/src/infra/db/postgres/models"
	"gorm.io/gorm"
)

type userService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) repositories.UserRepositoryImpl {
	return &userService{db: db}
}

func (s *userService) Upsert(ctx context.Context, data any, tx *gorm.DB) error {
	payload, ok := data.(*entities.UserEntity)
	if !ok {
		return errors.New("data is not a user entity")
	}
	model := pgmodels.UserModel{}
	model.FromEntity(payload)

	if err := s.db.WithContext(ctx).Save(&model).Error; err != nil {
		return err
	}
	return nil
}

func (s *userService) Detail(ctx context.Context, id string, tx *gorm.DB) (any, error) {
	model := pgmodels.UserModel{}
	if err := s.db.WithContext(ctx).Where("id = ?", id).First(&model).Error; err != nil {
		return nil, err
	}
	return model.ToEntity(), nil
}
