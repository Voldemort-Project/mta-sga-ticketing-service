package entities

import (
	"github.com/google/uuid"
)

type UserEntity struct {
	id             uuid.UUID
	name           string
	email          *string
	phoneNumber    *string
	idCardNumber   *string
	organizationID *uuid.UUID
	divisionID     *uuid.UUID
	organization   *OrganizationEntity
	division       *DivisionEntity
	roles          []RoleEntity
}

func (e *UserEntity) GetID() uuid.UUID {
	return e.id
}

func (e *UserEntity) GetName() string {
	return e.name
}

func (e *UserEntity) GetEmail() *string {
	return e.email
}

func (e *UserEntity) GetPhoneNumber() *string {
	return e.phoneNumber
}

func (e *UserEntity) GetIdCardNumber() *string {
	return e.idCardNumber
}

func (e *UserEntity) GetOrganizationID() *uuid.UUID {
	return e.organizationID
}

func (e *UserEntity) GetDivisionID() *uuid.UUID {
	return e.divisionID
}

func (e *UserEntity) GetOrganization() *OrganizationEntity {
	return e.organization
}

func (e *UserEntity) GetDivision() *DivisionEntity {
	return e.division
}

func (e *UserEntity) GetRoles() []RoleEntity {
	return e.roles
}

func (e *UserEntity) SetID(id uuid.UUID) {
	e.id = id
}

func (e *UserEntity) SetName(name string) {
	e.name = name
}

func (e *UserEntity) SetEmail(email string) {
	e.email = &email
}

func (e *UserEntity) SetPhoneNumber(phoneNumber string) {
	e.phoneNumber = &phoneNumber
}

func (e *UserEntity) SetIdCardNumber(idCardNumber string) {
	e.idCardNumber = &idCardNumber
}

func (e *UserEntity) SetOrganizationID(organizationID *uuid.UUID) {
	e.organizationID = organizationID
}

func (e *UserEntity) SetDivisionID(divisionID *uuid.UUID) {
	e.divisionID = divisionID
}

func (e *UserEntity) SetOrganization(organization *OrganizationEntity) {
	e.organization = organization
}

func (e *UserEntity) SetDivision(division *DivisionEntity) {
	e.division = division
}

func (e *UserEntity) SetRoles(roles []RoleEntity) {
	e.roles = roles
}

func MakeUserEntity(
	id uuid.UUID,
	name string,
	email *string,
	phoneNumber *string,
	idCardNumber *string,
	organizationID *uuid.UUID,
	divisionID *uuid.UUID,
) *UserEntity {
	return &UserEntity{
		id:             id,
		name:           name,
		email:          email,
		phoneNumber:    phoneNumber,
		idCardNumber:   idCardNumber,
		organizationID: organizationID,
		divisionID:     divisionID,
	}
}

func NewUserEntity(
	name string,
	email *string,
	phoneNumber *string,
	idCardNumber *string,
	organizationID *uuid.UUID,
	divisionID *uuid.UUID,
) *UserEntity {
	return MakeUserEntity(
		uuid.New(),
		name,
		email,
		phoneNumber,
		idCardNumber,
		organizationID,
		divisionID,
	)
}
