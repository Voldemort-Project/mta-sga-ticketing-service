package repositories

import (
	"context"

	"gorm.io/gorm"
)

type BaseRepositoryImpl interface {
	Upsert(ctx context.Context, data any, tx *gorm.DB) error
	Detail(ctx context.Context, id string, tx *gorm.DB) (any, error)
}
