package messages_tcp_transport

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/punnch/cli-messanger/internal/core/domain"
)

// /msg <message>
func (h *MessagesTCPHandler) SendMessage(
	ctx context.Context,
	userID uuid.UUID,
	roomID uuid.UUID,
	args []string,
) (domain.Message, error) {
	if len(args) < 1 {
		return domain.Message{}, fmt.Errorf("usage: /msg <message>")
	}

	content := strings.Join(args, " ")

	message := domain.CreateMessage(
		userID,
		roomID,
		content,
	)

	message, err := h.messagesRepository.SendMessage(ctx, message)
	if err != nil {
		return domain.Message{}, fmt.Errorf("send message to service: %w", err)
	}

	return message, nil
}
