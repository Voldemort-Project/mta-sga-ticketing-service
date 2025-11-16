package handlers

import (
	"github.com/Heian28/go-utils/fiber/goresponse"
	dtorequest "github.com/Voldemort-Project/sga-service/src/app/dto/request"
	"github.com/Voldemort-Project/sga-service/src/app/usecases"
	"github.com/gofiber/fiber/v3"
)

type CheckinHandlerImpl interface {
	GuestRegistration(c fiber.Ctx) error
}

type checkinHandler struct {
	resClient goresponse.GoResponseClient
	ucGuest   usecases.CheckinUsecaseImpl
}

func NewGuestHandler(
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
