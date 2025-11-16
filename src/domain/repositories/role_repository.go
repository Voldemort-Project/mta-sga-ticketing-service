package repositories

import (
	"context"

	"github.com/Voldemort-Project/sga-service/src/domain/entities"
	"gorm.io/gorm"
)

type RoleRepositoryImpl interface {
	BaseRepositoryImpl
	GetRoleByName(ctx context.Context, name string, tx *gorm.DB) (*entities.RoleEntity, error)
}
