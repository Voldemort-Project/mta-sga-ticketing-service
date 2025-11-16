package pgmodels

import (
	"github.com/Voldemort-Project/sga-service/src/constants"
	"github.com/Voldemort-Project/sga-service/src/domain/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrganizationModel struct {
	BaseModelWithUUID
	Name    string  `gorm:"column:name;not null"`
	Address *string `gorm:"column:address"`
}

func (OrganizationModel) TableName() string {
	return constants.TableNameOrganization
}

func (m *OrganizationModel) BeforeCreate(tx *gorm.DB) error {
	if m.ID == uuid.Nil {
		m.ID = uuid.New()
	}
	return nil
}

func (m *OrganizationModel) ToEntity() *entities.OrganizationEntity {
	return entities.MakeOrganizationEntity(
		m.ID,
		m.Name,
		m.Address,
	)
}

func (m *OrganizationModel) FromEntity(entity *entities.OrganizationEntity) {
	m.ID = entity.GetID()
	m.Name = entity.GetName()
	m.Address = entity.GetAddress()
}
