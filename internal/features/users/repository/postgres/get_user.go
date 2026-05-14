package users_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	domain "github.com/punnch/cli-messanger/internal/core/domain"
	core_errors "github.com/punnch/cli-messanger/internal/core/errors"
	core_postgres_pool "github.com/punnch/cli-messanger/internal/core/repository/postgres/pool"
)

func (r *UsersRepository) GetUser(
	ctx context.Context,
	username string,
) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	SELECT id, username, password_hash, created_at
	FROM messanger.users
	WHERE username=$1
	`

	row := r.pool.QueryRow(ctx, query, username)

	var userModel UserModel
	if err := userModel.Scan(row); err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.User{}, fmt.Errorf(
				"user with username='%s': %w",
				username,
				core_errors.ErrNotFound,
			)
		}

		return domain.User{}, fmt.Errorf("scan error: %w", err)
	}

	userDomain := modelToDomain(userModel)

	return userDomain, nil
}
