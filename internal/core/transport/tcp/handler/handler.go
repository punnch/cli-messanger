package handler_tcp_transport

import (
	"context"
	"net"

	domain "github.com/punnch/cli-messanger/internal/core/domain"
	core_logger "github.com/punnch/cli-messanger/internal/core/logger"

	"github.com/google/uuid"
)

type Handler struct {
	log *core_logger.Logger
	hub Hub

	usersHandler    UsersTCPHandler
	roomsHandler    RoomsTCPHandler
	messagesHandler MessagesTCPHandler
}

type Hub interface {
	SendMessage(
		roomID string,
		payload []byte,
	)

	Register(
		roomID string,
		conn net.Conn,
	)

	Unregister(
		roomID string,
		conn net.Conn,
	)

	Run()
}

type UsersTCPHandler interface {
	Register(
		ctx context.Context,
		args []string,
	) (domain.User, error)

	Login(
		ctx context.Context,
		args []string,
	) (domain.User, error)
}

type RoomsTCPHandler interface {
	ListRooms(
		ctx context.Context,
	) ([]domain.Room, error)

	JoinRoom(
		ctx context.Context,
		userID uuid.UUID,
		args []string,
	) (domain.Room, error)
}

type MessagesTCPHandler interface {
	GetHistory(
		ctx context.Context,
		roomID uuid.UUID,
	) ([]domain.Message, error)

	SendMessage(
		ctx context.Context,
		userID uuid.UUID,
		roomID uuid.UUID,
		args []string,
	) (domain.Message, error)
}

func NewHandler(
	log *core_logger.Logger,
	hub Hub,
	usersHandler UsersTCPHandler,
	roomsHandler RoomsTCPHandler,
	messagesHandler MessagesTCPHandler,
) *Handler {
	return &Handler{
		log: log,
		hub: hub,

		usersHandler:    usersHandler,
		roomsHandler:    roomsHandler,
		messagesHandler: messagesHandler,
	}
}
