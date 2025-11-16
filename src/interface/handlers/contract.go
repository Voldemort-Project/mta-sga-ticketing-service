package handlers

import (
	"github.com/Heian28/go-utils/fiber/goresponse"
	"github.com/Voldemort-Project/sga-service/src/app/usecases"
)

type AppHandler struct {
	CheckinHandler CheckinHandlerImpl
}

func NewAppHandler(resClient goresponse.GoResponseClient, usecaseApp usecases.AppUsecase) AppHandler {
	return AppHandler{
		CheckinHandler: NewCheckinHandler(
			resClient,
			usecaseApp.CheckinUsecase,
		),
	}
}
