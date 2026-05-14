package rooms_tcp_transport

import (
	"context"

	"github.com/punnch/cli-messanger/internal/core/domain"

	"github.com/google/uuid"
)

type RoomsTCPHandler struct {
	roomsService RoomsService
}

type RoomsService interface {
	ListRooms(
		ctx context.Context,
	) ([]domain.Room, error)

	JoinRoom(
		ctx context.Context,
		userID uuid.UUID,
		roomName string,
	) (domain.Room, error)
}

func NewRoomsTCPHandler(
	roomsService RoomsService,
) *RoomsTCPHandler {
	return &RoomsTCPHandler{
		roomsService: roomsService,
	}
}
