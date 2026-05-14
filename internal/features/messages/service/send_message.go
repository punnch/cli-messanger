package messages_service

import (
	"context"
	"fmt"

	"github.com/punnch/cli-messanger/internal/core/domain"
)

func (s *MessagesService) SendMessage(
	ctx context.Context,
	message domain.Message,
) (domain.Message, error) {
	message, err := s.messagesRepository.SaveMessage(ctx, message)
	if err != nil {
		return domain.Message{}, fmt.Errorf("save message in repository: %w", err)
	}

	return message, nil
}
