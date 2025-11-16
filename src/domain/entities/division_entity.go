package entities

import "github.com/google/uuid"

type DivisionEntity struct {
	id             uuid.UUID
	name           string
	description    *string
	organizationID uuid.UUID
	organization   *OrganizationEntity
}

func (e *DivisionEntity) GetID() uuid.UUID {
	return e.id
}

func (e *DivisionEntity) GetName() string {
	return e.name
}

func (e *DivisionEntity) GetDescription() *string {
	return e.description
}

func (e *DivisionEntity) GetOrganizationID() uuid.UUID {
	return e.organizationID
}

func (e *DivisionEntity) GetOrganization() *OrganizationEntity {
	return e.organization
}

func (e *DivisionEntity) SetID(id uuid.UUID) {
	e.id = id
}

func (e *DivisionEntity) SetName(name string) {
	e.name = name
}

func (e *DivisionEntity) SetDescription(description string) {
	e.description = &description
}

func (e *DivisionEntity) SetOrganizationID(organizationID uuid.UUID) {
	e.organizationID = organizationID
}

func (e *DivisionEntity) SetOrganization(organization *OrganizationEntity) {
	e.organization = organization
}

func MakeDivisionEntity(
	id uuid.UUID,
	name string,
	description *string,
	organizationID uuid.UUID,
) *DivisionEntity {
	return &DivisionEntity{
		id:             id,
		name:           name,
		description:    description,
		organizationID: organizationID,
	}
}
