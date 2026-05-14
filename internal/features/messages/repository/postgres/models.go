package messages_postgres_repository

import (
	"time"

	"github.com/google/uuid"
	"github.com/punnch/cli-messanger/internal/core/domain"
	core_postgres_pool "github.com/punnch/cli-messanger/internal/core/repository/postgres/pool"
)

type MessageModel struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	RoomID    uuid.UUID
	Username  string
	Content   string
	CreatedAt time.Time
}

func modelToDomain(model MessageModel) domain.Message {
	return domain.NewMessage(
		model.ID,
		model.UserID,
		model.RoomID,
		model.Username,
		model.Content,
		model.CreatedAt,
	)
}

func modelsToDomains(models []MessageModel) []domain.Message {
	domains := make([]domain.Message, len(models))

	for i, model := range models {
		domains[i] = modelToDomain(model)
	}

	return domains
}

func (m *MessageModel) Scan(row core_postgres_pool.Row) error {
	return row.Scan(
		&m.ID,
		&m.UserID,
		&m.RoomID,
		&m.Username,
		&m.Content,
		&m.CreatedAt,
	)
}
