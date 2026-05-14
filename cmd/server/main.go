package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	core_config "github.com/punnch/cli-messanger/internal/core/config"
	core_logger "github.com/punnch/cli-messanger/internal/core/logger"
	core_pgx_pool "github.com/punnch/cli-messanger/internal/core/repository/postgres/pool/pgx"
	handler_tcp_transport "github.com/punnch/cli-messanger/internal/core/transport/tcp/handler"
	core_tcp_server "github.com/punnch/cli-messanger/internal/core/transport/tcp/server"
	hub "github.com/punnch/cli-messanger/internal/features/hub"
	messages_postgres_repository "github.com/punnch/cli-messanger/internal/features/messages/repository/postgres"
	messages_service "github.com/punnch/cli-messanger/internal/features/messages/service"
	messages_tcp_transport "github.com/punnch/cli-messanger/internal/features/messages/transport/tcp"
	rooms_postgres_repository "github.com/punnch/cli-messanger/internal/features/rooms/repository/postgres"
	rooms_service "github.com/punnch/cli-messanger/internal/features/rooms/service"
	rooms_tcp_transport "github.com/punnch/cli-messanger/internal/features/rooms/transport/tcp"
	users_postgres_repository "github.com/punnch/cli-messanger/internal/features/users/repository/postgres"
	users_service "github.com/punnch/cli-messanger/internal/features/users/service"
	users_tcp_transport "github.com/punnch/cli-messanger/internal/features/users/transport/tcp"

	"go.uber.org/zap"
)

func main() {
	cfg := core_config.NewConfigMust()
	time.Local = cfg.TimeZone

	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT, syscall.SIGTERM,
	)
	defer cancel()

	logger, err := core_logger.NewLogger(core_logger.NewConfigMust())
	if err != nil {
		fmt.Println("failed to init application logger:", err)
		os.Exit(1)
	}

	logger.Debug("application time zone", zap.Any("zone", time.Local))

	logger.Debug("initializing postgresql connection pool")
	postgresPool, err := core_pgx_pool.NewPool(
		ctx,
		core_pgx_pool.NewConfigMust(),
	)
	if err != nil {
		logger.Fatal("failed to init postgresql connection pool", zap.Error(err))
	}
	defer postgresPool.Close()

	logger.Debug("initializing feature", zap.String("feature", "users"))
	usersRepository := users_postgres_repository.NewUsersRepository(postgresPool)
	usersService := users_service.NewUsersService(usersRepository)
	usersTransportTCP := users_tcp_transport.NewUsersTCPHandler(usersService)

	logger.Debug("initializing feature", zap.String("feature", "rooms"))
	roomsRepository := rooms_postgres_repository.NewRoomsRepository(postgresPool)
	roomsService := rooms_service.NewRoomsService(roomsRepository)
	roomsTransportTCP := rooms_tcp_transport.NewRoomsTCPHandler(roomsService)

	logger.Debug("initializing feature", zap.String("feature", "messages"))
	messagesRepository := messages_postgres_repository.NewMessagesRepository(postgresPool)
	messagerService := messages_service.NewMessagesService(messagesRepository)
	messagerTransportTCP := messages_tcp_transport.NewMessagesTCPTransport(messagerService)

	logger.Debug("initializing feature", zap.String("feature", "hub"))
	hub := hub.NewHub()
	go hub.Run()

	logger.Debug("initializing connection handler")
	connHandler := handler_tcp_transport.NewHandler(
		logger,
		hub,
		usersTransportTCP,
		roomsTransportTCP,
		messagerTransportTCP,
	)

	logger.Debug("initializing TCP server")
	tcpServer := core_tcp_server.NewTCPServer(
		core_tcp_server.NewConfigMust(),
		logger,
		connHandler.Handle,
	)
	if err := tcpServer.Run(ctx); err != nil {
		logger.Error("TCP server run error", zap.Error(err))
	}
}
