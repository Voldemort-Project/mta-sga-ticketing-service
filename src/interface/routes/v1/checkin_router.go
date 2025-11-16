package v1_router

import (
	"github.com/Voldemort-Project/sga-service/src/interface/handlers"
	"github.com/gofiber/fiber/v3"
)

func CheckinRouter(app fiber.Router, hndl handlers.CheckinHandlerImpl) {
	r := app.Group("/checkins")
	r.Post("/registration", hndl.GuestRegistration)
	r.Get("/guest", hndl.CheckinGuest)
}
