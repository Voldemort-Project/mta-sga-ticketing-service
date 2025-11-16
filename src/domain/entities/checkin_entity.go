package entities

import (
	"time"

	"github.com/google/uuid"
)

type CheckinEntity struct {
	id           uuid.UUID
	orgID        *uuid.UUID
	userID       uuid.UUID
	roomID       uuid.UUID
	checkInTime  time.Time
	checkoutTime *time.Time
	guestName    string
	guestEmail   *string
	guestPhone   *string
	guestIDCard  *string
	user         *UserEntity
	room         *RoomEntity
	organization *OrganizationEntity
}

func (e *CheckinEntity) GetID() uuid.UUID {
	return e.id
}

func (e *CheckinEntity) GetOrgID() *uuid.UUID {
	return e.orgID
}

func (e *CheckinEntity) GetUserID() uuid.UUID {
	return e.userID
}

func (e *CheckinEntity) GetRoomID() uuid.UUID {
	return e.roomID
}

func (e *CheckinEntity) GetCheckInTime() time.Time {
	return e.checkInTime
}

func (e *CheckinEntity) GetCheckoutTime() *time.Time {
	return e.checkoutTime
}

func (e *CheckinEntity) GetGuestName() string {
	return e.guestName
}

func (e *CheckinEntity) GetGuestEmail() *string {
	return e.guestEmail
}

func (e *CheckinEntity) GetGuestPhone() *string {
	return e.guestPhone
}

func (e *CheckinEntity) GetGuestIDCard() *string {
	return e.guestIDCard
}

func (e *CheckinEntity) GetUser() *UserEntity {
	return e.user
}

func (e *CheckinEntity) GetRoom() *RoomEntity {
	return e.room
}

func (e *CheckinEntity) GetOrganization() *OrganizationEntity {
	return e.organization
}

func (e *CheckinEntity) SetID(id uuid.UUID) {
	e.id = id
}

func (e *CheckinEntity) SetOrgID(orgID *uuid.UUID) {
	e.orgID = orgID
}

func (e *CheckinEntity) SetUserID(userID uuid.UUID) {
	e.userID = userID
}

func (e *CheckinEntity) SetRoomID(roomID uuid.UUID) {
	e.roomID = roomID
}

func (e *CheckinEntity) SetCheckInTime(checkInTime time.Time) {
	e.checkInTime = checkInTime
}

func (e *CheckinEntity) SetCheckoutTime(checkoutTime *time.Time) {
	e.checkoutTime = checkoutTime
}

func (e *CheckinEntity) SetGuestName(guestName string) {
	e.guestName = guestName
}

func (e *CheckinEntity) SetGuestEmail(guestEmail string) {
	e.guestEmail = &guestEmail
}

func (e *CheckinEntity) SetGuestPhone(guestPhone string) {
	e.guestPhone = &guestPhone
}

func (e *CheckinEntity) SetGuestIDCard(guestIDCard string) {
	e.guestIDCard = &guestIDCard
}

func (e *CheckinEntity) SetUser(user *UserEntity) {
	e.user = user
}

func (e *CheckinEntity) SetRoom(room *RoomEntity) {
	e.room = room
}

func (e *CheckinEntity) SetOrganization(organization *OrganizationEntity) {
	e.organization = organization
}

func MakeCheckinEntity(
	id uuid.UUID,
	orgID *uuid.UUID,
	userID uuid.UUID,
	roomID uuid.UUID,
	checkInTime time.Time,
	checkoutTime *time.Time,
	guestName string,
	guestEmail *string,
	guestPhone *string,
	guestIDCard *string,
) *CheckinEntity {
	return &CheckinEntity{
		id:           id,
		orgID:        orgID,
		userID:       userID,
		roomID:       roomID,
		checkInTime:  checkInTime,
		checkoutTime: checkoutTime,
		guestName:    guestName,
		guestEmail:   guestEmail,
		guestPhone:   guestPhone,
		guestIDCard:  guestIDCard,
	}
}

func NewCheckinEntity(
	orgID *uuid.UUID,
	userID uuid.UUID,
	roomID uuid.UUID,
	checkInTime time.Time,
	checkoutTime *time.Time,
	guestName string,
	guestEmail *string,
	guestPhone *string,
	guestIDCard *string,
) *CheckinEntity {
	return MakeCheckinEntity(
		uuid.New(),
		orgID,
		userID,
		roomID,
		checkInTime,
		checkoutTime,
		guestName,
		guestEmail,
		guestPhone,
		guestIDCard,
	)
}
