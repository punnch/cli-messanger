package users_tcp_transport

import (
	"context"

	"github.com/punnch/cli-messanger/internal/core/domain"
)

type UsersTCPHandler struct {
	usersService UsersService
}

type UsersService interface {
	Register(
		ctx context.Context,
		username string,
		password string,
	) (domain.User, error)

	Login(
		ctx context.Context,
		username string,
		password string,
	) (domain.User, error)
}

func NewUsersTCPHandler(
	usersService UsersService,
) *UsersTCPHandler {
	return &UsersTCPHandler{
		usersService: usersService,
	}
}
