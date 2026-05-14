package domain

import "github.com/google/uuid"

type Member struct {
	UserID uuid.UUID
	RoomID uuid.UUID
}

func NewMember(
	userID uuid.UUID,
	roomID uuid.UUID,
) Member {
	return Member{
		UserID: userID,
		RoomID: roomID,
	}
}
