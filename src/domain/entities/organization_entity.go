package entities

import "github.com/google/uuid"

type OrganizationEntity struct {
	id      uuid.UUID
	name    string
	address *string
}

func (e *OrganizationEntity) GetID() uuid.UUID {
	return e.id
}

func (e *OrganizationEntity) GetName() string {
	return e.name
}

func (e *OrganizationEntity) GetAddress() *string {
	return e.address
}

func (e *OrganizationEntity) SetID(id uuid.UUID) {
	e.id = id
}

func (e *OrganizationEntity) SetName(name string) {
	e.name = name
}

func (e *OrganizationEntity) SetAddress(address string) {
	e.address = &address
}

func MakeOrganizationEntity(
	id uuid.UUID,
	name string,
	address *string,
) *OrganizationEntity {
	return &OrganizationEntity{
		id:      id,
		name:    name,
		address: address,
	}
}
