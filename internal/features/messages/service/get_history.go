package messages_service

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/punnch/cli-messanger/internal/core/domain"
)

func (s *MessagesService) GetHistory(
	ctx context.Context,
	roomID uuid.UUID,
) ([]domain.Message, error) {
	limit := 50

	messages, err := s.messagesRepository.GetMessages(ctx, roomID, &limit)
	if err != nil {
		return nil, fmt.Errorf("get messages from repository: %w", err)
	}

	return messages, nil
}
