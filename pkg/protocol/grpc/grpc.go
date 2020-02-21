package grpc

import (
	"context"
	v1 "github.com/soulmonk/cuppa-workers-authentication/pkg/api/v1"
	"github.com/soulmonk/cuppa-workers-authentication/pkg/logger"
	"github.com/soulmonk/cuppa-workers-authentication/pkg/protocol/grpc/middleware"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
)

// RunServer runs HTTP/REST gateway
func RunServer(ctx context.Context, v1API v1.AuthenticationServiceServer, port string) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	// gRPC server statup options
	var opts []grpc.ServerOption

	// add middleware
	opts = middleware.AddLogging(logger.Log, opts)

	// register service
	server := grpc.NewServer(opts...)
	v1.RegisterAuthenticationServiceServer(server, v1API)

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			logger.Log.Warn("shutting down gRPC server...")

			server.GracefulStop()

			<-ctx.Done()
		}
	}()

	// start gRPC server
	logger.Log.Info("starting gRPC server... at :" + port)
	return server.Serve(listen)
}
