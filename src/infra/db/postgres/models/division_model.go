package pgmodels

import (
	"github.com/Voldemort-Project/sga-service/src/constants"
	"github.com/Voldemort-Project/sga-service/src/domain/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DivisionModel struct {
	BaseModelWithUUID
	OrganizationID uuid.UUID          `gorm:"column:organization_id;not null;"`
	Name           string             `gorm:"column:name;not null"`
	Description    *string            `gorm:"column:description"`
	Organization   *OrganizationModel `gorm:"foreignKey:OrganizationID;save_associations:false"`
}

func (DivisionModel) TableName() string {
	return constants.TableNameDivision
}

func (m *DivisionModel) BeforeCreate(tx *gorm.DB) error {
	if m.ID == uuid.Nil {
		m.ID = uuid.New()
	}
	return nil
}

func (m *DivisionModel) ToEntity() *entities.DivisionEntity {
	data := entities.MakeDivisionEntity(
		m.ID,
		m.Name,
		m.Description,
		m.OrganizationID,
	)

	if m.Organization != nil {
		org := m.Organization.ToEntity()
		data.SetOrganization(org)
	}

	return data
}

func (m *DivisionModel) FromEntity(entity *entities.DivisionEntity) {
	m.ID = entity.GetID()
	m.Name = entity.GetName()
	m.Description = entity.GetDescription()
	m.OrganizationID = entity.GetOrganizationID()

	if entity.GetOrganization() != nil {
		m.Organization = &OrganizationModel{}
		m.Organization.FromEntity(entity.GetOrganization())
	}
}
