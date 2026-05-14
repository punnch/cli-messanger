package users_tcp_transport

import (
	"context"
	"fmt"
	"strings"

	"github.com/punnch/cli-messanger/internal/core/domain"
)

func (h *UsersTCPHandler) Register(
	ctx context.Context,
	args []string,
) (domain.User, error) {
	if len(args) != 2 {
		return domain.User{}, fmt.Errorf("usage: /register <username> <password>")
	}

	username := strings.TrimSpace(args[0])
	password := strings.TrimSpace(args[1])

	user, err := h.usersService.Register(
		ctx,
		username,
		password,
	)
	if err != nil {
		return domain.User{}, fmt.Errorf(
			"register user: %s: %w",
			username,
			err,
		)
	}

	return user, nil
}
