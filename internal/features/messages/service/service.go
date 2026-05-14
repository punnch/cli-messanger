package messages_service

import (
	"context"

	"github.com/google/uuid"

	"github.com/punnch/cli-messanger/internal/core/domain"
)

type MessagesService struct {
	messagesRepository MessagesRepository
}

type MessagesRepository interface {
	GetMessages(
		ctx context.Context,
		roomID uuid.UUID,
		limit *int,
	) ([]domain.Message, error)

	SaveMessage(
		ctx context.Context,
		message domain.Message,
	) (domain.Message, error)
}

func NewMessagesService(
	messagesRepository MessagesRepository,
) *MessagesService {
	return &MessagesService{
		messagesRepository: messagesRepository,
	}
}
