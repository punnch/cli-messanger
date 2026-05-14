package users_service

import (
	"context"
	"fmt"

	"github.com/punnch/cli-messanger/internal/core/domain"
	"golang.org/x/crypto/bcrypt"
)

func (s *UsersService) Login(
	ctx context.Context,
	username string,
	password string,
) (domain.User, error) {
	user, err := s.usersRepository.GetUser(ctx, username)
	if err != nil {
		return domain.User{}, fmt.Errorf("get user from repository: %w", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return domain.User{}, fmt.Errorf("password is not correct : %w", err)
	}

	return user, nil
}
