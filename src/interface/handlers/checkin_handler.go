package handlers

import (
	"github.com/Heian28/go-utils/fiber/goresponse"
	dtorequest "github.com/Voldemort-Project/sga-service/src/app/dto/request"
	"github.com/Voldemort-Project/sga-service/src/app/usecases"
	"github.com/Voldemort-Project/sga-service/utils"
	"github.com/gofiber/fiber/v3"
)

type CheckinHandlerImpl interface {
	GuestRegistration(c fiber.Ctx) error
	CheckinGuest(c fiber.Ctx) error
}

type checkinHandler struct {
	resClient goresponse.GoResponseClient
	ucGuest   usecases.CheckinUsecaseImpl
}

func NewCheckinHandler(
	resClient goresponse.GoResponseClient,
	ucGuest usecases.CheckinUsecaseImpl,
) CheckinHandlerImpl {
	return &checkinHandler{
		resClient: resClient,
		ucGuest:   ucGuest,
	}
}

func (h *checkinHandler) GuestRegistration(c fiber.Ctx) error {
	ctx := c.Context()
	dto, err := dtorequest.NewCheckinRegistrationRequestDto(c)
	if err != nil {
		panic(err)
	}
	if err := h.ucGuest.RegistrationGuest(ctx, dto); err != nil {
		panic(err)
	}
	return h.resClient.Jsonify(
		c,
		fiber.StatusCreated,
		"Checkin registration successful",
		nil,
		nil,
	)
}

func (h *checkinHandler) CheckinGuest(c fiber.Ctx) error {
	ctx := c.Context()
	pagination := utils.NewPagination(c)
	rows, total, err := h.ucGuest.GetCheckinGuestList(ctx, pagination)
	if err != nil {
		panic(err)
	}
	meta := h.resClient.CreateMeta(pagination.Page, pagination.PerPage, int(total))
	return h.resClient.Jsonify(
		c,
		fiber.StatusOK,
		"Fetch checkin guest list successfully",
		rows,
		meta,
	)
}
