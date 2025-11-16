package router

import (
	"github.com/Voldemort-Project/sga-service/src/interface/handlers"
	v1_router "github.com/Voldemort-Project/sga-service/src/interface/routes/v1"
	"github.com/gofiber/fiber/v3"
)

func InitRouter(app *fiber.App, hndl handlers.AppHandler) {
	appRouter := app.Group("/api")

	// Define router using versioning
	v1_router.InitV1Router(appRouter, hndl)
}
