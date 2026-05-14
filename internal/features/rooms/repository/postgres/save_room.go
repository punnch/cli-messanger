package rooms_postgres_repository

import (
	"context"
	"fmt"

	"github.com/punnch/cli-messanger/internal/core/domain"
)

func (r *RoomsRepository) SaveRoom(
	ctx context.Context,
	room domain.Room,
) (domain.Room, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	INSERT INTO messanger.rooms (id, name, created_at)
	VALUES ($1, $2, $3)
	RETURNING id, name, created_at;
	`

	row := r.pool.QueryRow(
		ctx,
		query,
		room.ID,
		room.Name,
		room.CreatedAt,
	)

	var roomModel RoomModel
	if err := roomModel.Scan(row); err != nil {
		return domain.Room{}, fmt.Errorf("scan error: %w", err)
	}

	roomDomain := roomModelToDomain(roomModel)

	return roomDomain, nil
}
