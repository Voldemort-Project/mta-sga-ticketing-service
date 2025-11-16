package v1_router

import (
	"github.com/Voldemort-Project/sga-service/src/interface/handlers"
	"github.com/gofiber/fiber/v3"
)

func InitV1Router(app fiber.Router, hndlr handlers.AppHandler) {
	router := app.Group("/v1")
	CheckinRouter(router, hndlr.CheckinHandler)
}
