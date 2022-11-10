package grpc

import (
	"context"
	admin_v1 "github.com/soulmonk/cuppa-workers-authentication/pkg/api/admin"
	authentication_v1 "github.com/soulmonk/cuppa-workers-authentication/pkg/api/authentication"
	"github.com/soulmonk/cuppa-workers-authentication/pkg/db"
	"github.com/soulmonk/cuppa-workers-authentication/pkg/logger"
	"github.com/soulmonk/cuppa-workers-authentication/pkg/protocol/grpc/middleware"
	admin_service_v1 "github.com/soulmonk/cuppa-workers-authentication/pkg/service/admin/v1"
	authentication_service_v1 "github.com/soulmonk/cuppa-workers-authentication/pkg/service/authentication/v1"
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

	// gRPC server statup options
	var opts []grpc.ServerOption

	// add middleware
	opts = middleware.AddLogging(logger.Log, opts)

	// register service
	server := grpc.NewServer(opts...)
	//Register reflection service on gRPC server. to allow use describe command
	reflection.Register(server)
	authenticationV1API := authentication_service_v1.NewAuthenticationServiceServer(dao)
	authentication_v1.RegisterAuthenticationServiceServer(server, authenticationV1API)
	adminV1API := admin_service_v1.NewAuthenticationServiceServer(dao)
	admin_v1.RegisterAdminServiceServer(server, adminV1API)

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
