package usecases

import pgservice "github.com/Voldemort-Project/sga-service/src/infra/db/postgres/service"

type AppUsecase struct {
	CheckinUsecase CheckinUsecaseImpl
}

func NewAppUsecase(srv pgservice.AppPostgreService) AppUsecase {
	return AppUsecase{
		CheckinUsecase: NewCheckinUsecase(
			srv.UserService,
			srv.RoomService,
			srv.CheckinService,
			srv.RoleService,
			srv.UserRoleService,
		),
	}
}
