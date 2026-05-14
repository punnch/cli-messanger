package rooms_postgres_repository

import (
	"context"
	"fmt"

	"github.com/punnch/cli-messanger/internal/core/domain"
)

func (r *RoomsRepository) SaveMember(
	ctx context.Context,
	member domain.Member,
) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	INSERT INTO messanger.room_members (user_id, room_id)
	VALUES ($1, $2)
	ON CONFLICT DO NOTHING
	`

	_, err := r.pool.Exec(
		ctx,
		query,
		member.UserID,
		member.RoomID,
	)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}
