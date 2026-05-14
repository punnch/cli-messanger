package domain

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	core_errors "github.com/punnch/cli-messanger/internal/core/errors"
)

type Room struct {
	ID        uuid.UUID
	Name      string
	CreatedAt time.Time
}

func NewRoom(
	id uuid.UUID,
	name string,
	createdAt time.Time,
) Room {
	return Room{
		ID:        id,
		Name:      name,
		CreatedAt: createdAt,
	}
}

func CreateRoom(name string) Room {
	var (
		id        = uuid.New()
		createdAt = time.Now()
	)

	return Room{
		ID:        id,
		Name:      name,
		CreatedAt: createdAt,
	}
}

func (r *Room) Validate() error {
	roomNameLen := len([]rune(r.Name))
	if roomNameLen < 1 || roomNameLen > 50 {
		return fmt.Errorf(
			"invalid 'RoomName' len: %d: %w",
			roomNameLen,
			core_errors.ErrInvalidArgument,
		)
	}

	return nil
}

func RoomsPrettyFormat(rooms []Room) string {
	var sb strings.Builder

	for i, room := range rooms {
		fmt.Fprintf(&sb, "%d. %s\n", i+1, room.Name)
	}

	return sb.String()
}
