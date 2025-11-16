package pgmodels

import (
	"github.com/Voldemort-Project/sga-service/src/constants"
	"github.com/Voldemort-Project/sga-service/src/domain/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RoleModel struct {
	BaseModelWithUUID
	Name        string  `gorm:"column:name;not null"`
	Description *string `gorm:"column:description"`
}

func (RoleModel) TableName() string {
	return constants.TableNameRole
}

func (m *RoleModel) BeforeCreate(tx *gorm.DB) error {
	if m.ID == uuid.Nil {
		m.ID = uuid.New()
	}
	return nil
}

func (m *RoleModel) ToEntity() *entities.RoleEntity {
	return entities.MakeRoleEntity(
		m.ID,
		m.Name,
		m.Description,
	)
}

func (m *RoleModel) FromEntity(entity *entities.RoleEntity) {
	m.ID = entity.GetID()
	m.Name = entity.GetName()
	m.Description = entity.GetDescription()
}
