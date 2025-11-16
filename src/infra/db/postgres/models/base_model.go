package pgmodels

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseParanoidTimetamp struct {
	CreatedAt time.Time       `gorm:"column:created_at;not null;default:now();type:timestamp with time zone"`
	UpdatedAt time.Time       `gorm:"column:updated_at;not null;default:now();type:timestamp with time zone"`
	DeletedAt *gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp with time zone"`
}

type BaseModelWithSerialID struct {
	BaseParanoidTimetamp
	ID uint `gorm:"primaryKey;autoIncrement:true;column:id"`
}

type BaseModelWithUUID struct {
	BaseParanoidTimetamp
	ID uuid.UUID `gorm:"primaryKey;column:id"`
}
