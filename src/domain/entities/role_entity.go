package entities

import "github.com/google/uuid"

type RoleEntity struct {
	id          uuid.UUID
	name        string
	description *string
}

func (e *RoleEntity) GetID() uuid.UUID {
	return e.id
}

func (e *RoleEntity) GetName() string {
	return e.name
}

func (e *RoleEntity) GetDescription() *string {
	return e.description
}

func (e *RoleEntity) SetID(id uuid.UUID) {
	e.id = id
}

func (e *RoleEntity) SetName(name string) {
	e.name = name
}

func (e *RoleEntity) SetDescription(description string) {
	e.description = &description
}

func MakeRoleEntity(
	id uuid.UUID,
	name string,
	description *string,
) *RoleEntity {
	return &RoleEntity{
		id:          id,
		name:        name,
		description: description,
	}
}
