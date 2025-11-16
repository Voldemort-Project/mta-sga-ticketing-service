package v1_router

import (
	"github.com/Voldemort-Project/sga-service/src/interface/handlers"
	"github.com/gofiber/fiber/v3"
)

func CheckinRouter(app fiber.Router, hndl handlers.CheckinHandlerImpl) {
	rtr := app.Group("/checkins")
	rtr.Post("/registration", hndl.GuestRegistration)
}
