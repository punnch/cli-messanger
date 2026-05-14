package rooms_service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/punnch/cli-messanger/internal/core/domain"
)

func (s *RoomsService) JoinRoom(
	ctx context.Context,
	userID uuid.UUID,
	roomName string,
) (domain.Room, error) {
	room, err := s.roomsRepository.GetRoom(ctx, roomName)
	if err != nil {
		room = domain.CreateRoom(roomName)

		if err := room.Validate(); err != nil {
			return domain.Room{}, fmt.Errorf("validate room domain: %w", err)
		}

		room, err = s.roomsRepository.SaveRoom(ctx, room)
		if err != nil {
			return domain.Room{}, fmt.Errorf("save room in repository: %w", err)
		}
	}

	member := domain.NewMember(userID, room.ID)

	if err := s.roomsRepository.SaveMember(ctx, member); err != nil {
		return domain.Room{}, fmt.Errorf("save member in repository: %w", err)
	}

	return room, nil
}
