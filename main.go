package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Heian28/go-utils/db/gopostgres"
	"github.com/Heian28/go-utils/fiber/goerror"
	"github.com/Heian28/go-utils/fiber/gomiddleware"
	"github.com/Heian28/go-utils/fiber/goresponse"
	"github.com/Heian28/go-utils/gologger"
	"github.com/Voldemort-Project/sga-service/configs"
	"github.com/Voldemort-Project/sga-service/src/app/usecases"
	pgservice "github.com/Voldemort-Project/sga-service/src/infra/db/postgres/service"
	infraerror "github.com/Voldemort-Project/sga-service/src/infra/error"
	"github.com/Voldemort-Project/sga-service/src/interface/handlers"
	router "github.com/Voldemort-Project/sga-service/src/interface/routes"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/sirupsen/logrus"
)

func main() {
	configs.Init()
	var conf configs.Config = configs.AppConfig

	isProduction := conf.App.Env == "production"

	// Initialize the logger
	gologger.New(
		gologger.SetIsProduction(isProduction),
		gologger.SetServiceName(conf.App.Name),
	)
	log := gologger.Logger

	goerror.RegisterGoFiberError(
		infraerror.RegisterAppError(),
		infraerror.RegisterHttpError(),
	)

	log.Infof("%+v", conf.Postgres)

	// Initialize the database
	postgres := gopostgres.New(
		isProduction,
		conf.Postgres,
		log,
	)
	defer postgres.Close()

	// Define postgres transaction
	db := postgres.Database()
	gopostgres.NewGoPostgresTransaction(db)

	app := fiber.New(newFiberConfig(conf, log))

	app.Use(recover.New())
	app.Use(gomiddleware.RequestID())
	app.Use(gomiddleware.HttpLogger(log))

	// Initialize response client
	resClient := goresponse.NewGoResponseClient()

	// Initialize the postgres service
	postgresService := pgservice.NewAppPostgreService(db)
	// Initialize the usecases
	usecaseApp := usecases.NewAppUsecase(postgresService)
	// Initialize the handlers
	hndl := handlers.NewAppHandler(resClient, usecaseApp)

	// Initialize the router app
	router.InitRouter(app, hndl)

	log.Fatal(app.Listen(
		fmt.Sprintf(":%d", conf.App.Port),
		fiber.ListenConfig{
			EnablePrintRoutes: true,
			GracefulContext:   context.Background(),
			ShutdownTimeout:   5 * time.Second,
		},
	))
}

func newFiberConfig(conf configs.Config, log *logrus.Logger) fiber.Config {
	return fiber.Config{
		AppName:      conf.App.Name,
		BodyLimit:    1024 * 1024 * 10, // 10MB
		ErrorHandler: goerror.NewErrorHandler(log),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
}
