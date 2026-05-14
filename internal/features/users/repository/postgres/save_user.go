package users_postgres_repository

import (
	"context"
	"fmt"

	"github.com/punnch/cli-messanger/internal/core/domain"
)

func (r *UsersRepository) SaveUser(
	ctx context.Context,
	user domain.User,
) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	INSERT INTO messanger.users (id, username, password_hash, created_at)
	VALUES ($1, $2, $3, $4)
	RETURNING id, username, password_hash, created_at;
	`

	row := r.pool.QueryRow(
		ctx,
		query,
		user.ID,
		user.Username,
		user.PasswordHash,
		user.CreatedAt,
	)

	var userModel UserModel
	if err := userModel.Scan(row); err != nil {
		return domain.User{}, fmt.Errorf("scan error: %w", err)
	}

	userDomain := modelToDomain(userModel)

	return userDomain, nil
}
