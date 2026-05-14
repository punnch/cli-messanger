package rooms_postgres_repository

import (
	"context"
	"fmt"

	"github.com/punnch/cli-messanger/internal/core/domain"
)

func (r *RoomsRepository) GetRoom(
	ctx context.Context,
	name string,
) (domain.Room, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	SELECT id, name, created_at 
	FROM messanger.rooms
	WHERE name=$1;
	`

	row := r.pool.QueryRow(
		ctx,
		query,
		name,
	)

	var roomModel RoomModel
	if err := roomModel.Scan(row); err != nil {
		return domain.Room{}, fmt.Errorf("scan error: %w", err)
	}

	roomDomains := roomModelToDomain(roomModel)

	return roomDomains, nil
}
