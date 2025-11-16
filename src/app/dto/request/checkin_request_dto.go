package dtorequest

import (
	"regexp"
	"time"

	"github.com/Heian28/go-utils/fiber/goerror"
	infraerror "github.com/Voldemort-Project/sga-service/src/infra/error"
	"github.com/Voldemort-Project/sga-service/utils"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/gofiber/fiber/v3"
)

type CheckinRegistrationRequestDto struct {
	Name         string  `json:"name"`
	Email        string  `json:"email"`
	PhoneNumber  *string `json:"phoneNumber"`
	IDCardNumber *string `json:"idCardNumber"`
	RoomNumber   string  `json:"roomNumber"`
	CheckinDate  string  `json:"checkinDate"`
}

func (dto *CheckinRegistrationRequestDto) Validate() error {
	return validation.ValidateStruct(dto,
		validation.Field(&dto.Name, validation.Required),
		validation.Field(
			&dto.Email,
			validation.Required,
			is.Email,
		),
		validation.Field(
			&dto.PhoneNumber,
			validation.When(dto.PhoneNumber != nil, validation.Match(regexp.MustCompile(`^[0-9]{10,15}$`))),
		),
		validation.Field(
			&dto.RoomNumber,
			validation.Required,
			validation.Match(regexp.MustCompile(`^[0-9]{3,15}$`)),
		),
		validation.Field(&dto.CheckinDate, validation.Required, validation.By(func(value any) error {
			s, _ := value.(string)

			d, err := time.Parse(time.DateTime, s)
			if err != nil {
				return validation.NewError("checkinDate", "must be a valid date format with YYYY-MM-DD HH:mm:ss")
			}
			if d.Before(time.Now()) {
				return validation.NewError("checkinDate", "must be in the future")
			}

			return nil
		})),
	)
}

func NewCheckinRegistrationRequestDto(c fiber.Ctx) (*CheckinRegistrationRequestDto, error) {
	dto := &CheckinRegistrationRequestDto{}
	if err := utils.CopyJSON(dto, c.Body()); err != nil {
		ne := goerror.ComposeClientError(infraerror.ErrBadRequestInvalidParseBody, err)
		ne.SetServerMessage(err.Error())
		return nil, ne
	}
	if err := dto.Validate(); err != nil {
		ne := goerror.ComposeClientError(infraerror.ErrUnprocessableEntity, err)
		ne.SetClientMessage(err.Error())
		return nil, ne
	}
	return dto, nil
}
