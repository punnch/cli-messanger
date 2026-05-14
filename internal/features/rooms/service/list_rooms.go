package rooms_service

import (
	"context"
	"fmt"

	"github.com/punnch/cli-messanger/internal/core/domain"
)

func (s *RoomsService) ListRooms(
	ctx context.Context,
) ([]domain.Room, error) {
	rooms, err := s.roomsRepository.GetRooms(ctx)
	if err != nil {
		return nil, fmt.Errorf("get rooms from repository: %w", err)
	}

	return rooms, nil
}
