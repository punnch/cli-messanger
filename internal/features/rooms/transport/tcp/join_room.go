package rooms_tcp_transport

import (
	"context"
	"fmt"
	"strings"

	"github.com/punnch/cli-messanger/internal/core/domain"

	"github.com/google/uuid"
)

// /join <room>.
func (h *RoomsTCPHandler) JoinRoom(
	ctx context.Context,
	userID uuid.UUID,
	args []string,
) (domain.Room, error) {
	if len(args) != 1 {
		return domain.Room{}, fmt.Errorf("usage: /join <room>")
	}

	roomName := strings.TrimSpace(args[0])

	room, err := h.roomsService.JoinRoom(ctx, userID, roomName)
	if err != nil {
		return domain.Room{}, fmt.Errorf("join room: %w", err)
	}

	return room, nil
}
