package pgservice

import (
	"github.com/Voldemort-Project/sga-service/src/domain/repositories"
	"gorm.io/gorm"
)

type AppPostgreService struct {
	CheckinService  repositories.CheckinRepositoryImpl
	RoleService     repositories.RoleRepositoryImpl
	RoomService     repositories.RoomRepositoryImpl
	UserService     repositories.UserRepositoryImpl
	UserRoleService repositories.UserRoleRepositoryImpl
}

func NewAppPostgreService(db *gorm.DB) AppPostgreService {
	return AppPostgreService{
		CheckinService:  NewCheckinService(db),
		RoleService:     NewRoleService(db),
		RoomService:     NewRoomService(db),
		UserService:     NewUserService(db),
		UserRoleService: NewUserRoleService(db),
	}
}
