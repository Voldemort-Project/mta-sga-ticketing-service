package entities

import "github.com/google/uuid"

type UserRoleEntity struct {
	userID uuid.UUID
	roleID uuid.UUID
	user   *UserEntity
	role   *RoleEntity
}

func (e *UserRoleEntity) GetUserID() uuid.UUID {
	return e.userID
}

func (e *UserRoleEntity) GetRoleID() uuid.UUID {
	return e.roleID
}

func (e *UserRoleEntity) GetUser() *UserEntity {
	return e.user
}

func (e *UserRoleEntity) GetRole() *RoleEntity {
	return e.role
}

func (e *UserRoleEntity) SetUserID(userID uuid.UUID) {
	e.userID = userID
}

func (e *UserRoleEntity) SetRoleID(roleID uuid.UUID) {
	e.roleID = roleID
}

func (e *UserRoleEntity) SetUser(user *UserEntity) {
	e.user = user
}

func (e *UserRoleEntity) SetRole(role *RoleEntity) {
	e.role = role
}

func MakeUserRoleEntity(
	userID uuid.UUID,
	roleID uuid.UUID,
) *UserRoleEntity {
	return &UserRoleEntity{
		userID: userID,
		roleID: roleID,
	}
}

func NewUserRoleEntity() *UserRoleEntity {
	return &UserRoleEntity{}
}
