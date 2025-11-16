package infraerror

import (
	"github.com/Heian28/go-utils/fiber/goerror"
)

const (
	ErrInternalServerError        goerror.GoFiberErrorCode = 500_001_00000
	ErrUnprocessableEntity        goerror.GoFiberErrorCode = 422_001_00000
	ErrBadRequestInvalidParseBody goerror.GoFiberErrorCode = 400_001_00000

	ErrResourceNotFound     goerror.GoFiberErrorCode = 404_001_00000
	ErrResourceRoomNotFound goerror.GoFiberErrorCode = 404_001_00001

	ErrResourceRoomNotAvailable goerror.GoFiberErrorCode = 400_001_00001
)

func RegisterAppError() map[goerror.GoFiberErrorCode]*goerror.GoFiberErrorCommon {
	return map[goerror.GoFiberErrorCode]*goerror.GoFiberErrorCommon{
		// Default errors
		ErrInternalServerError: {
			ErrorCode:     ErrInternalServerError,
			ClientMessage: "An error has occurred. Please try again later.",
			ServerMessage: "Internal Server Error",
		},
		ErrUnprocessableEntity: {
			ErrorCode:     ErrUnprocessableEntity,
			ClientMessage: "The request was well-formed but was unable to be followed due to semantic errors.",
			ServerMessage: "Unprocessable Entity",
		},
		ErrBadRequestInvalidParseBody: {
			ErrorCode:     ErrBadRequestInvalidParseBody,
			ClientMessage: "The request body is invalid. Please check the request body and try again.",
			ServerMessage: "Bad Request Invalid Parse Body",
		},

		// Not found errors
		ErrResourceNotFound: {
			ErrorCode:     ErrResourceNotFound,
			ClientMessage: "The resource was not found.",
			ServerMessage: "Resource Not Found",
		},
		ErrResourceRoomNotFound: {
			ErrorCode:     ErrResourceRoomNotFound,
			ClientMessage: "The room was not found.",
			ServerMessage: "Room Not Found",
		},
		ErrResourceRoomNotAvailable: {
			ErrorCode:     ErrResourceRoomNotAvailable,
			ClientMessage: "The room is not available.",
			ServerMessage: "Room Not Available",
		},
	}
}
