package users_service

import (
	"context"
	"fmt"

	"github.com/punnch/cli-messanger/internal/core/domain"
	"golang.org/x/crypto/bcrypt"
)

func (s *UsersService) Register(
	ctx context.Context,
	username string,
	password string,
) (domain.User, error) {
	user := domain.CreateUser(
		username,
		password,
	)

	if err := user.Validate(); err != nil {
		return domain.User{}, fmt.Errorf("validate user domain: %w", err)
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return domain.User{}, fmt.Errorf("hash password: %w", err)
	}

	user.PasswordHash = string(hash)

	user, err = s.usersRepository.SaveUser(ctx, user)
	if err != nil {
		return domain.User{}, fmt.Errorf("save user in repository: %w", err)
	}

	return user, nil
}
