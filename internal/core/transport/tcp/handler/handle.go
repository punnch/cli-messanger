package handler_tcp_transport

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"strings"

	domain "github.com/punnch/cli-messanger/internal/core/domain"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

func (h *Handler) Handle(
	ctx context.Context,
	conn net.Conn,
) {
	defer conn.Close()

	var (
		session       *domain.User
		currentRoomID string
	)

	defer func() {
		if currentRoomID != "" {
			h.hub.Unregister(currentRoomID, conn)
		}
	}()

	write := func(msg string) {
		conn.Write([]byte(msg))
	}

	writeln := func(msg string) {
		write(msg + "\n")
	}

	writeln("Welcome! Use /register <username> <password> or /login <username> <password>\n")
	write("(/help) " + help())

	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if errors.Is(err, io.EOF) {
				h.log.Info("client disconnected", zap.String("addr", conn.RemoteAddr().String()))
				return
			}

			h.log.Error("read buf: %w", zap.Error(err))
			return
		}

		line := strings.TrimSpace(string(buf[:n]))
		if line == "" {
			continue
		}

		fields := strings.Fields(line)

		cmd := fields[0]
		args := fields[1:]

		switch cmd {

		case "/help":
			write(help())

		case "/register":
			if session != nil {
				writeln("ERR " + "you already have an account")
				continue
			}

			user, err := h.usersHandler.Register(ctx, args)
			if err != nil {
				h.log.Warn("register failed", zap.Error(err))
				writeln("ERR " + err.Error())
				continue
			}

			session = &user
			writeln("OK registered and logged in as " + user.Username)

		case "/login":
			if session != nil {
				writeln("ERR " + "you already have an account")
				continue
			}

			user, err := h.usersHandler.Login(ctx, args)
			if err != nil {
				h.log.Warn("login failed", zap.Error(err))
				writeln("ERR " + err.Error())
				continue
			}

			session = &user
			writeln("OK logged in as " + user.Username)

		case "/join":
			if session == nil {
				writeln("ERR not authenticated - use /register or /login (/help for guide)")
				continue
			}

			room, err := h.roomsHandler.JoinRoom(ctx, session.ID, args)
			if err != nil {
				h.log.Warn("failed to join in room", zap.Error(err))
				writeln("ERR " + err.Error())
				continue
			}

			if currentRoomID != "" {
				h.hub.Unregister(currentRoomID, conn)
			}

			currentRoomID = room.ID.String()

			h.hub.Register(currentRoomID, conn)

			writeln("OK joined in room " + room.Name)

		case "/rooms":
			rooms, err := h.roomsHandler.ListRooms(ctx)
			if err != nil {
				h.log.Error("failed to list rooms", zap.Error(err))
				writeln("ERR " + err.Error())
				continue
			}

			roomList := domain.RoomsPrettyFormat(rooms)

			write(roomList)

		case "/history":
			if session == nil {
				writeln("ERR not authenticated - use /register or /login (/help for guide)")
				continue
			}

			if currentRoomID == "" {
				writeln("ERR not in room - use /join")
				continue
			}

			messages, err := h.messagesHandler.GetHistory(
				ctx,
				uuid.MustParse(currentRoomID),
			)
			if err != nil {
				h.log.Error("failed to get message from room", zap.Error(err))
				writeln("ERR " + err.Error())
				continue
			}

			messageList := domain.MessagesPrettyFormat(messages)

			write(messageList)

		case "/msg":
			if session == nil {
				writeln("ERR not authenticated - use /register or /login (/help for guide)")
				continue
			}

			if currentRoomID == "" {
				writeln("ERR not in a room - use /join")
				continue
			}

			msg, err := h.messagesHandler.SendMessage(
				ctx,
				session.ID,
				uuid.MustParse(currentRoomID),
				args,
			)
			if err != nil {
				h.log.Error("failed to send message", zap.Error(err))
				writeln("ERR " + err.Error())
				continue
			}

			payload := fmt.Sprintf(
				"[%s] %s\n",
				session.Username,
				msg.Content,
			)

			h.hub.SendMessage(currentRoomID, []byte(payload))

		default:
			writeln("ERR unknown command: " + cmd)
		}
	}
}
