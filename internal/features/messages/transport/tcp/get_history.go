package messages_tcp_transport

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/punnch/cli-messanger/internal/core/domain"
)

func (h *MessagesTCPHandler) GetHistory(
	ctx context.Context,
	roomID uuid.UUID,
) ([]domain.Message, error) {
	messages, err := h.messagesRepository.GetHistory(ctx, roomID)
	if err != nil {
		return nil, fmt.Errorf("get history from service: %w", err)
	}

	return messages, nil
}
