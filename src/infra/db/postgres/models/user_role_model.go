package pgmodels

import (
	"github.com/Voldemort-Project/sga-service/src/constants"
	"github.com/Voldemort-Project/sga-service/src/domain/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRoleModel struct {
	BaseModelWithUUID
	UserID uuid.UUID  `gorm:"column:user_id;not null"`
	RoleID uuid.UUID  `gorm:"column:role_id;not null"`
	User   *UserModel `gorm:"foreignKey:UserID;save_associations:true"`
	Role   *RoleModel `gorm:"foreignKey:RoleID;save_associations:true"`
}

func (UserRoleModel) TableName() string {
	return constants.TableNameUserRole
}

func (m *UserRoleModel) BeforeCreate(tx *gorm.DB) error {
	if m.ID == uuid.Nil {
		m.ID = uuid.New()
	}
	return nil
}

func (m *UserRoleModel) ToEntity() *entities.UserRoleEntity {
	return entities.MakeUserRoleEntity(
		m.UserID,
		m.RoleID,
	)
}

func (m *UserRoleModel) FromEntity(entity *entities.UserRoleEntity) {
	m.UserID = entity.GetUserID()
	m.RoleID = entity.GetRoleID()

	if entity.GetUser() != nil {
		m.User = &UserModel{}
		m.User.FromEntity(entity.GetUser())
	}
	if entity.GetRole() != nil {
		m.Role = &RoleModel{}
		m.Role.FromEntity(entity.GetRole())
	}
}
