package rooms_service

import (
	"context"

	"github.com/punnch/cli-messanger/internal/core/domain"
)

type RoomsService struct {
	roomsRepository RoomsRepository
}

type RoomsRepository interface {
	GetRooms(
		ctx context.Context,
	) ([]domain.Room, error)

	GetRoom(
		ctx context.Context,
		name string,
	) (domain.Room, error)

	SaveRoom(
		ctx context.Context,
		room domain.Room,
	) (domain.Room, error)

	SaveMember(
		ctx context.Context,
		member domain.Member,
	) error
}

func NewRoomsService(
	roomsRepository RoomsRepository,
) *RoomsService {
	return &RoomsService{
		roomsRepository: roomsRepository,
	}
}
