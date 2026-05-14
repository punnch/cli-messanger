package messages_postgres_repository

import (
	"context"
	"fmt"

	"github.com/punnch/cli-messanger/internal/core/domain"
)

func (r *MessagesRepository) SaveMessage(
	ctx context.Context,
	message domain.Message,
) (domain.Message, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	INSERT INTO messanger.messages (id, user_id, room_id, content, created_at)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id, user_id, room_id, content, created_at;
	`

	_, err := r.pool.Exec(
		ctx,
		query,
		message.ID,
		message.UserID,
		message.RoomID,
		message.Content,
		message.CreatedAt,
	)
	if err != nil {
		return domain.Message{}, fmt.Errorf("exec error: %w", err)
	}

	return message, nil
}
