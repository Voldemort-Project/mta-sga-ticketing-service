package usecases

import (
	"context"
	"errors"
	"time"

	"github.com/Heian28/go-utils/db/gopostgres"
	"github.com/Heian28/go-utils/fiber/goerror"
	dtorequest "github.com/Voldemort-Project/sga-service/src/app/dto/request"
	"github.com/Voldemort-Project/sga-service/src/domain/entities"
	"github.com/Voldemort-Project/sga-service/src/domain/repositories"
	infraerror "github.com/Voldemort-Project/sga-service/src/infra/error"
	"github.com/Voldemort-Project/sga-service/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CheckinUsecaseImpl interface {
	RegistrationGuest(
		ctx context.Context,
		dto *dtorequest.CheckinRegistrationRequestDto,
	) error
	GetCheckinGuestList(
		ctx context.Context,
		pagination *utils.PaginationDto,
	) ([]entities.CheckinEntity, int64, error)
}

type checkinUsecase struct {
	userRepo     repositories.UserRepositoryImpl
	roomRepo     repositories.RoomRepositoryImpl
	checkinRepo  repositories.CheckinRepositoryImpl
	roleRepo     repositories.RoleRepositoryImpl
	userRoleRepo repositories.UserRoleRepositoryImpl
}

func NewCheckinUsecase(
	userRepo repositories.UserRepositoryImpl,
	roomRepo repositories.RoomRepositoryImpl,
	checkinRepo repositories.CheckinRepositoryImpl,
	roleRepo repositories.RoleRepositoryImpl,
	userRoleRepo repositories.UserRoleRepositoryImpl,
) CheckinUsecaseImpl {
	return &checkinUsecase{
		userRepo:     userRepo,
		roomRepo:     roomRepo,
		checkinRepo:  checkinRepo,
		roleRepo:     roleRepo,
		userRoleRepo: userRoleRepo,
	}
}

func (u *checkinUsecase) RegistrationGuest(
	ctx context.Context,
	dto *dtorequest.CheckinRegistrationRequestDto,
) error {
	checkinDate, _ := time.Parse(time.DateTime, dto.CheckinDate)
	var TRX = gopostgres.GoPostgresTransaction

	if err := TRX.WithTransaction(ctx, func(tx *gorm.DB) error {
		room, err := u.roomRepo.GetRoomByNameAndAvailibility(ctx, dto.RoomNumber, tx)
		if err != nil {
			ne := goerror.ComposeClientError(infraerror.ErrResourceRoomNotFound, err)
			return ne
		}
		if !room.GetIsAvailable() {
			ne := goerror.ComposeClientError(infraerror.ErrResourceRoomNotAvailable, err)
			return ne
		}
		role, _ := u.roleRepo.GetRoleByName(ctx, "guest", tx)
		if role == nil {
			return goerror.ComposeClientError(infraerror.ErrResourceNotFound, errors.New("role not found"))
		}

		userPayload := entities.NewUserEntity(
			dto.Name,
			&dto.Email,
			dto.PhoneNumber,
			dto.IDCardNumber,
			nil,
			nil,
		)

		checkinPayload := entities.NewCheckinEntity(
			nil,
			uuid.Nil,
			room.GetID(),
			checkinDate,
			nil,
			dto.Name,
			&dto.Email,
			dto.PhoneNumber,
			dto.IDCardNumber,
		)

		userRolePayload := entities.NewUserRoleEntity()

		checkinPayload.SetUser(userPayload)
		userRolePayload.SetUser(userPayload)
		userRolePayload.SetRole(role)

		if err := u.checkinRepo.Upsert(ctx, checkinPayload, tx); err != nil {
			ne := goerror.ComposeClientError(infraerror.ErrInternalServerError, err)
			ne.SetServerMessage(err.Error())
			return ne
		}

		if err := u.userRoleRepo.Upsert(ctx, userRolePayload, tx); err != nil {
			ne := goerror.ComposeClientError(infraerror.ErrInternalServerError, err)
			ne.SetServerMessage(err.Error())
			return ne
		}

		if err := u.roomRepo.UpdateVisibility(ctx, room.GetID().String(), false, tx); err != nil {
			ne := goerror.ComposeClientError(infraerror.ErrInternalServerError, err)
			ne.SetServerMessage(err.Error())
			return ne
		}

		return nil
	}); err != nil {
		return err
	}
	return nil
}

func (u *checkinUsecase) GetCheckinGuestList(
	ctx context.Context,
	pagination *utils.PaginationDto,
) ([]entities.CheckinEntity, int64, error) {
	rows, total, err := u.checkinRepo.GetCheckinGuestList(ctx, pagination)
	if err != nil {
		ne := goerror.ComposeClientError(infraerror.ErrInternalServerError, err)
		ne.SetServerMessage(err.Error())
		return nil, 0, ne
	}
	return rows, total, nil
}
