package entities

import "github.com/google/uuid"

type RoomEntity struct {
	id          uuid.UUID
	name        string
	description *string
	isAvailable bool
}

func (e *RoomEntity) GetID() uuid.UUID {
	return e.id
}

func (e *RoomEntity) GetName() string {
	return e.name
}

func (e *RoomEntity) GetDescription() *string {
	return e.description
}

func (e *RoomEntity) GetIsAvailable() bool {
	return e.isAvailable
}

func (e *RoomEntity) SetID(id uuid.UUID) {
	e.id = id
}

func (e *RoomEntity) SetName(name string) {
	e.name = name
}

func (e *RoomEntity) SetDescription(description string) {
	e.description = &description
}

func (e *RoomEntity) SetIsAvailable(isAvailable bool) {
	e.isAvailable = isAvailable
}

func MakeRoomEntity(
	id uuid.UUID,
	name string,
	description *string,
	isAvailable bool,
) *RoomEntity {
	return &RoomEntity{
		id:          id,
		name:        name,
		description: description,
		isAvailable: isAvailable,
	}
}
