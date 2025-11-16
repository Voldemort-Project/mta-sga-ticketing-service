package dtoresponse

import (
	"time"

	"github.com/Voldemort-Project/sga-service/src/domain/entities"
)

type CheckinGuestResponseDto struct {
	ID           string  `json:"id"`
	OrgID        *string `json:"orgId"`
	UserID       string  `json:"userId"`
	RoomID       string  `json:"roomId"`
	CheckInTime  string  `json:"checkInTime"`
	CheckOutTime *string `json:"checkOutTime"`
	GuestName    string  `json:"guestName"`
	GuestEmail   *string `json:"guestEmail"`
	GuestPhone   *string `json:"guestPhone"`
	GuestIDCard  *string `json:"guestIdCard"`
}

func (dto *CheckinGuestResponseDto) FromEntity(entity *entities.CheckinEntity) {
	dto.ID = entity.GetID().String()
	if entity.GetOrgID() != nil {
		orgID := entity.GetOrgID().String()
		dto.OrgID = &orgID
	}
	dto.UserID = entity.GetUserID().String()
	dto.RoomID = entity.GetRoomID().String()
	dto.CheckInTime = entity.GetCheckInTime().Format(time.DateTime)
	if entity.GetCheckoutTime() != nil {
		checkOutTime := entity.GetCheckoutTime().Format(time.DateTime)
		dto.CheckOutTime = &checkOutTime
	}
	dto.GuestName = entity.GetGuestName()
	dto.GuestEmail = entity.GetGuestEmail()
	dto.GuestPhone = entity.GetGuestPhone()
	dto.GuestIDCard = entity.GetGuestIDCard()
}

func NewCheckinGuestResponseDto() *CheckinGuestResponseDto {
	return &CheckinGuestResponseDto{}
}

func TransformListCheckinGuestResponseDto(rows []entities.CheckinEntity) []CheckinGuestResponseDto {
	response := make([]CheckinGuestResponseDto, 0)
	for _, row := range rows {
		d := NewCheckinGuestResponseDto()
		d.FromEntity(&row)
		response = append(response, *d)
	}
	return response
}
