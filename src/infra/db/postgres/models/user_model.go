package pgmodels

import (
	"github.com/Voldemort-Project/sga-service/src/constants"
	"github.com/Voldemort-Project/sga-service/src/domain/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserModel struct {
	BaseModelWithUUID
	Name           string             `gorm:"column:name;not null"`
	Email          *string            `gorm:"column:email;"`
	PhoneNumber    *string            `gorm:"column:phone_number"`
	IDCardNumber   *string            `gorm:"column:id_card_number"`
	OrganizationID *uuid.UUID         `gorm:"column:organization_id;constraint:OnDelete:SET NULL;constraint:OnUpdate:SET NULL"`
	DivisionID     *uuid.UUID         `gorm:"column:division_id;constraint:OnDelete:SET NULL;constraint:OnUpdate:SET NULL"`
	Organization   *OrganizationModel `gorm:"foreignKey:OrganizationID;save_associations:false"`
	Division       *DivisionModel     `gorm:"foreignKey:DivisionID;save_associations:false"`
	Roles          []RoleModel        `gorm:"many2many:user_roles;save_associations:true"`
}

func (UserModel) TableName() string {
	return constants.TableNameUser
}

func (m *UserModel) BeforeCreate(tx *gorm.DB) error {
	if m.ID == uuid.Nil {
		m.ID = uuid.New()
	}
	return nil
}

func (m *UserModel) ToEntity() *entities.UserEntity {
	data := entities.MakeUserEntity(
		m.ID,
		m.Name,
		m.Email,
		m.PhoneNumber,
		m.IDCardNumber,
		m.OrganizationID,
		m.DivisionID,
	)
	if m.Organization != nil {
		org := m.Organization.ToEntity()
		data.SetOrganization(org)
	}
	if m.Division != nil {
		division := m.Division.ToEntity()
		data.SetDivision(division)
	}
	if m.Roles != nil {
		roles := make([]entities.RoleEntity, len(m.Roles))
		for i, role := range m.Roles {
			roleData := role.ToEntity()
			roles[i] = *roleData
		}
		data.SetRoles(roles)
	}
	return data
}

func (m *UserModel) FromEntity(entity *entities.UserEntity) {
	m.ID = entity.GetID()
	m.Name = entity.GetName()
	m.Email = entity.GetEmail()
	m.PhoneNumber = entity.GetPhoneNumber()
	m.IDCardNumber = entity.GetIdCardNumber()
	m.OrganizationID = entity.GetOrganizationID()
	m.DivisionID = entity.GetDivisionID()

	if entity.GetOrganization() != nil {
		m.Organization = &OrganizationModel{}
		m.Organization.FromEntity(entity.GetOrganization())
	}
	if entity.GetDivision() != nil {
		m.Division = &DivisionModel{}
		m.Division.FromEntity(entity.GetDivision())
	}
	if entity.GetRoles() != nil {
		roles := make([]RoleModel, len(entity.GetRoles()))
		for i, role := range entity.GetRoles() {
			roles[i].FromEntity(&role)
		}
		m.Roles = roles
	}
}
