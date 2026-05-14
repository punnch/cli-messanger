package rooms_postgres_repository

import (
	"time"

	"github.com/google/uuid"
	"github.com/punnch/cli-messanger/internal/core/domain"
	core_postgres_pool "github.com/punnch/cli-messanger/internal/core/repository/postgres/pool"
)

// Room
type RoomModel struct {
	ID        uuid.UUID
	Name      string
	CreatedAt time.Time
}

func roomModelToDomain(model RoomModel) domain.Room {
	return domain.NewRoom(
		model.ID,
		model.Name,
		model.CreatedAt,
	)
}

func modelsToDomains(models []RoomModel) []domain.Room {
	domains := make([]domain.Room, len(models))

	for i, model := range models {
		domains[i] = roomModelToDomain(model)
	}

	return domains
}

func (m *RoomModel) Scan(row core_postgres_pool.Row) error {
	return row.Scan(
		&m.ID,
		&m.Name,
		&m.CreatedAt,
	)
}

// Member
type MemberModel struct {
	UserID uuid.UUID
	RoomID uuid.UUID
}

func memberModelToDomain(model MemberModel) domain.Member {
	return domain.NewMember(
		model.UserID,
		model.RoomID,
	)
}

func (m *MemberModel) Scan(row core_postgres_pool.Row) error {
	return row.Scan(
		&m.UserID,
		&m.RoomID,
	)
}
