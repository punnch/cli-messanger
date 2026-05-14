package messages_tcp_transport

import (
	"context"

	"github.com/google/uuid"

	"github.com/punnch/cli-messanger/internal/core/domain"
)

type MessagesTCPHandler struct {
	messagesRepository MessagesRepository
}

type MessagesRepository interface {
	SendMessage(
		ctx context.Context,
		message domain.Message,
	) (domain.Message, error)

	GetHistory(
		ctx context.Context,
		roomID uuid.UUID,
	) ([]domain.Message, error)
}

func NewMessagesTCPTransport(
	messagesRepository MessagesRepository,
) *MessagesTCPHandler {
	return &MessagesTCPHandler{
		messagesRepository: messagesRepository,
	}
}
