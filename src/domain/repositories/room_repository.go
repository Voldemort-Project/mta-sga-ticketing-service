package repositories

import (
	"context"

	"github.com/Voldemort-Project/sga-service/src/domain/entities"
	"gorm.io/gorm"
)

type RoomRepositoryImpl interface {
	BaseRepositoryImpl
	GetRoomByNameAndAvailibility(ctx context.Context, name string, tx *gorm.DB) (*entities.RoomEntity, error)
	UpdateVisibility(ctx context.Context, id string, isAvailable bool, tx *gorm.DB) error
}
