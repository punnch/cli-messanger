package domain

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	core_errors "github.com/punnch/cli-messanger/internal/core/errors"
)

type Message struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	RoomID    uuid.UUID
	Username  string
	Content   string
	CreatedAt time.Time
}

func NewMessage(
	id uuid.UUID,
	userID uuid.UUID,
	roomID uuid.UUID,
	username string,
	content string,
	createdAt time.Time,
) Message {
	return Message{
		ID:        id,
		UserID:    userID,
		RoomID:    roomID,
		Username:  username,
		Content:   content,
		CreatedAt: createdAt,
	}
}

func CreateMessage(
	userID uuid.UUID,
	roomID uuid.UUID,
	content string,
) Message {
	var (
		id        = uuid.New()
		createdAt = time.Now()
	)

	return Message{
		ID:        id,
		UserID:    userID,
		RoomID:    roomID,
		Content:   content,
		CreatedAt: createdAt,
	}
}

func (m *Message) Validate() error {
	messageContentLen := len([]rune(m.Content))
	if messageContentLen < 1 || messageContentLen > 2_000 {
		return fmt.Errorf(
			"invalid 'MessageContent' len: %d: %w",
			messageContentLen,
			core_errors.ErrInvalidArgument,
		)
	}

	return nil
}

func MessagesPrettyFormat(messages []Message) string {
	var sb strings.Builder

	for _, msg := range messages {
		fmt.Fprintf(&sb, "[%s] %s\n", msg.Username, msg.Content)
	}

	return sb.String()
}
