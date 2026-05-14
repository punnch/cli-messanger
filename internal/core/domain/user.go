package domain

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	core_errors "github.com/punnch/cli-messanger/internal/core/errors"
)

type User struct {
	ID           uuid.UUID
	Username     string
	PasswordHash string
	CreatedAt    time.Time
}

func NewUser(
	id uuid.UUID,
	username string,
	passwordHash string,
	createdAt time.Time,
) User {
	return User{
		ID:           id,
		Username:     username,
		PasswordHash: passwordHash,
		CreatedAt:    createdAt,
	}
}

func CreateUser(
	username string,
	passwordHash string,
) User {
	var (
		id        = uuid.New()
		createdAt = time.Now()
	)

	return User{
		ID:           id,
		Username:     username,
		PasswordHash: passwordHash,
		CreatedAt:    createdAt,
	}
}

func (u *User) Validate() error {
	usernameLen := len([]rune(u.Username))
	if usernameLen < 1 || usernameLen > 16 {
		return fmt.Errorf(
			"invalid 'UserName' len: %d: %w",
			usernameLen,
			core_errors.ErrInvalidArgument,
		)
	}

	passwordHashLen := len([]rune(u.PasswordHash))
	if passwordHashLen < 8 || passwordHashLen > 32 {
		return fmt.Errorf(
			"invalid 'PasswordHash len: %d: %w",
			passwordHashLen,
			core_errors.ErrInvalidArgument,
		)
	}

	return nil
}
