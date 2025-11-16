package pgmodels

import (
	"github.com/Voldemort-Project/sga-service/src/constants"
	"github.com/Voldemort-Project/sga-service/src/domain/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RoomModel struct {
	BaseModelWithUUID
	Name        string  `gorm:"column:name;not null"`
	Description *string `gorm:"column:description"`
	IsAvailable bool    `gorm:"column:is_available;not null;default:true"`
}

func (RoomModel) TableName() string {
	return constants.TableNameRoom
}

func (m *RoomModel) BeforeCreate(tx *gorm.DB) error {
	if m.ID == uuid.Nil {
		m.ID = uuid.New()
	}
	return nil
}

func (m *RoomModel) ToEntity() *entities.RoomEntity {
	return entities.MakeRoomEntity(
		m.ID,
		m.Name,
		m.Description,
		m.IsAvailable,
	)
}

func (m *RoomModel) FromEntity(entity *entities.RoomEntity) {
	m.ID = entity.GetID()
	m.Name = entity.GetName()
	m.Description = entity.GetDescription()
	m.IsAvailable = entity.GetIsAvailable()
}
