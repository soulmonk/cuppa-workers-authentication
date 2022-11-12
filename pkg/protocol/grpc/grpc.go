package grpc

import (
	"context"
	"github.com/soulmonk/cuppa-workers-authentication/pkg/api/admin"
	"github.com/soulmonk/cuppa-workers-authentication/pkg/api/authentication"
	"github.com/soulmonk/cuppa-workers-authentication/pkg/db"
	"github.com/soulmonk/cuppa-workers-authentication/pkg/logger"
	"github.com/soulmonk/cuppa-workers-authentication/pkg/protocol/grpc/middleware"
	adminService "github.com/soulmonk/cuppa-workers-authentication/pkg/service/admin"
	authenticationService "github.com/soulmonk/cuppa-workers-authentication/pkg/service/authentication"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
	"os/signal"
)

// RunServer runs HTTP/REST gateway
func RunServer(ctx context.Context, dao *db.Dao, port string) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	// gRPC server startup options
	var opts []grpc.ServerOption

	// add middleware
	opts = middleware.AddLogging(logger.Log, opts)

	// register service
	server := grpc.NewServer(opts...)
	//Register reflection service on gRPC server. to allow use describe command
	reflection.Register(server)
	authenticationAPI := authenticationService.NewAuthenticationServiceServer(dao)
	authentication.RegisterAuthenticationServiceServer(server, authenticationAPI)
	adminAPI := adminService.NewAuthenticationServiceServer(dao)
	admin.RegisterAdminServiceServer(server, adminAPI)

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
