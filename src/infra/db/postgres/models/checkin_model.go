package pgmodels

import (
	"time"

	"github.com/Voldemort-Project/sga-service/src/constants"
	"github.com/Voldemort-Project/sga-service/src/domain/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CheckinModel struct {
	BaseModelWithUUID
	OrgID        *uuid.UUID         `gorm:"column:org_id;type:text;default:null"`
	UserID       uuid.UUID          `gorm:"column:user_id;type:text;default:null"`
	RoomID       uuid.UUID          `gorm:"column:room_id;type:text;default:null"`
	CheckInTime  time.Time          `gorm:"column:check_in_time;type:timestamp with time zone;not null"`
	CheckOutTime *time.Time         `gorm:"column:check_out_time;type:timestamp with time zone;default:null"`
	GuestName    string             `gorm:"column:guest_name;not null"`
	GuestEmail   *string            `gorm:"column:guest_email;type:text;default:null"`
	GuestPhone   *string            `gorm:"column:guest_phone;type:text;default:null"`
	GuestIDCard  *string            `gorm:"column:guest_id_card;type:text;default:null"`
	Organization *OrganizationModel `gorm:"foreignKey:OrgID;save_associations:false;constraint:OnDelete:SET NULL;constraint:OnUpdate:SET NULL"`
	User         *UserModel         `gorm:"foreignKey:UserID;save_associations:true;constraint:OnDelete:SET NULL;constraint:OnUpdate:SET NULL"`
	Room         *RoomModel         `gorm:"foreignKey:RoomID;save_associations:false;constraint:OnDelete:SET NULL;constraint:OnUpdate:SET NULL"`
}

func (CheckinModel) TableName() string {
	return constants.TableNameCheckin
}

func (m *CheckinModel) BeforeCreate(tx *gorm.DB) error {
	if m.ID == uuid.Nil {
		m.ID = uuid.New()
	}
	return nil
}

func (m *CheckinModel) ToEntity() *entities.CheckinEntity {
	data := entities.MakeCheckinEntity(
		m.ID,
		m.OrgID,
		m.UserID,
		m.RoomID,
		m.CheckInTime,
		m.CheckOutTime,
		m.GuestName,
		m.GuestEmail,
		m.GuestPhone,
		m.GuestIDCard,
	)
	if m.User != nil {
		user := m.User.ToEntity()
		data.SetUser(user)
	}
	if m.Room != nil {
		room := m.Room.ToEntity()
		data.SetRoom(room)
	}
	if m.Organization != nil {
		organization := m.Organization.ToEntity()
		data.SetOrganization(organization)
	}
	return data
}

func (m *CheckinModel) FromEntity(entity *entities.CheckinEntity) {
	m.ID = entity.GetID()
	m.OrgID = entity.GetOrgID()
	m.UserID = entity.GetUserID()
	m.RoomID = entity.GetRoomID()
	m.CheckInTime = entity.GetCheckInTime()
	m.CheckOutTime = entity.GetCheckoutTime()
	m.GuestName = entity.GetGuestName()
	m.GuestEmail = entity.GetGuestEmail()
	m.GuestPhone = entity.GetGuestPhone()
	m.GuestIDCard = entity.GetGuestIDCard()

	if entity.GetUser() != nil {
		m.User = &UserModel{}
		m.User.FromEntity(entity.GetUser())
	}
}
