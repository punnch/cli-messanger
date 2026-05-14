package rooms_postgres_repository

import (
	"context"
	"fmt"

	"github.com/punnch/cli-messanger/internal/core/domain"
)

func (r *RoomsRepository) GetRooms(
	ctx context.Context,
) ([]domain.Room, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	SELECT id, name, created_at 
	FROM messanger.rooms
	ORDER BY id ASC;
	`

	rows, err := r.pool.Query(
		ctx,
		query,
	)
	if err != nil {
		return nil, fmt.Errorf("select rooms: %w", err)
	}
	defer rows.Close()

	var roomModels []RoomModel
	for rows.Next() {
		var roomModel RoomModel
		if err := roomModel.Scan(rows); err != nil {
			return nil, fmt.Errorf("scan error: %w", err)
		}

		roomModels = append(roomModels, roomModel)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("next rows: %w", err)
	}

	roomDomains := modelsToDomains(roomModels)

	return roomDomains, nil
}
