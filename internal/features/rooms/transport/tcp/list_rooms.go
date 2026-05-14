package rooms_tcp_transport

import (
	"context"
	"fmt"

	"github.com/punnch/cli-messanger/internal/core/domain"
)

// /rooms
func (h *RoomsTCPHandler) ListRooms(
	ctx context.Context,
) ([]domain.Room, error) {
	rooms, err := h.roomsService.ListRooms(ctx)
	if err != nil {
		return nil, fmt.Errorf("get rooms from service: %w", err)
	}

	return rooms, nil
}
