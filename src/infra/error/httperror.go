package infraerror

import (
	"github.com/Heian28/go-utils/fiber/goerror"
	"github.com/gofiber/fiber/v3"
)

func RegisterHttpError() map[goerror.GoFiberErrorCode]int {
	return map[goerror.GoFiberErrorCode]int{
		ErrInternalServerError:        fiber.StatusInternalServerError,
		ErrUnprocessableEntity:        fiber.StatusUnprocessableEntity,
		ErrBadRequestInvalidParseBody: fiber.StatusBadRequest,

		ErrResourceNotFound:     fiber.StatusNotFound,
		ErrResourceRoomNotFound: fiber.StatusNotFound,
	}
}
