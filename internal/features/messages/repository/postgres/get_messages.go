package messages_postgres_repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/punnch/cli-messanger/internal/core/domain"
)

func (r *MessagesRepository) GetMessages(
	ctx context.Context,
	roomID uuid.UUID,
	limit *int,
) ([]domain.Message, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	SELECT 
		m.id,
		m.user_id,
		m.room_id,
		u.username,
		m.content,
		m.created_at
	FROM messanger.messages m

	JOIN messanger.users u ON u.id = m.user_id

	WHERE m.room_id = $1

	ORDER BY m.created_at ASC
	LIMIT $2;
	`

	rows, err := r.pool.Query(
		ctx,
		query,
		roomID,
		limit,
	)
	if err != nil {
		return nil, fmt.Errorf("select messages: %w", err)
	}
	defer rows.Close()

	var messageModels []MessageModel
	for rows.Next() {
		var messageModel MessageModel
		if err := messageModel.Scan(rows); err != nil {
			return nil, fmt.Errorf("scan error: %w", err)
		}

		messageModels = append(messageModels, messageModel)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("next rows: %w", err)
	}

	messageDomains := modelsToDomains(messageModels)

	return messageDomains, nil

}
