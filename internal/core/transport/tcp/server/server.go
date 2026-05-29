package core_tcp_server

import (
	"context"
	"fmt"
	"net"

	core_logger "github.com/punnch/cli-messanger/internal/core/logger"

	"go.uber.org/zap"
)

type ConnHandler func(ctx context.Context, conn net.Conn)

type TCPServer struct {
	config  Config
	log     *core_logger.Logger
	handler ConnHandler
}

func NewTCPServer(
	cfg Config,
	log *core_logger.Logger,
	handler ConnHandler,
) *TCPServer {
	return &TCPServer{
		config:  cfg,
		log:     log,
		handler: handler,
	}
}

func (s *TCPServer) Run(ctx context.Context) error {
	ln, err := net.Listen("tcp", s.config.Addr)
	if err != nil {
		return fmt.Errorf("listen TCP server: %w", err)
	}

	s.log.Info("start TCP server", zap.String("addr", s.config.Addr))

	go func() {
		<-ctx.Done()
		if err := ln.Close(); err != nil {
			s.log.Error("close connection", zap.Error(err))
		}
	}()

	for {
		conn, err := ln.Accept()
		if err != nil {
			select {
			case <-ctx.Done():
				s.log.Warn("TCP server closed")
				return nil
			default:
				s.log.Error("TCP accept", zap.Error(err))
				continue
			}
		}

		go s.handler(ctx, conn)
	}
}
