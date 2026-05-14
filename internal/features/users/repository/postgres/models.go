package users_postgres_repository

import (
	"time"

	"github.com/google/uuid"
	"github.com/punnch/cli-messanger/internal/core/domain"
	core_postgres_pool "github.com/punnch/cli-messanger/internal/core/repository/postgres/pool"
)

type UserModel struct {
	ID           uuid.UUID
	Username     string
	PasswordHash string
	CreatedAt    time.Time
}

func modelToDomain(model UserModel) domain.User {
	return domain.User{
		ID:           model.ID,
		Username:     model.Username,
		PasswordHash: model.PasswordHash,
		CreatedAt:    model.CreatedAt,
	}
}

func (m *UserModel) Scan(row core_postgres_pool.Row) error {
	return row.Scan(
		&m.ID,
		&m.Username,
		&m.PasswordHash,
		&m.CreatedAt,
	)
}
