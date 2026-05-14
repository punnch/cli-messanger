package users_service

import (
	"context"

	"github.com/punnch/cli-messanger/internal/core/domain"
)

type UsersService struct {
	usersRepository UsersRepository
}

type UsersRepository interface {
	GetUser(
		ctx context.Context,
		username string,
	) (domain.User, error)

	SaveUser(
		ctx context.Context,
		user domain.User,
	) (domain.User, error)
}

func NewUsersService(
	usersRepository UsersRepository,
) *UsersService {
	return &UsersService{
		usersRepository: usersRepository,
	}
}
