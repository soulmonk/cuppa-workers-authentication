package cmd

import (
	"context"
	"fmt"
	"github.com/soulmonk/cuppa-workers-authentication/pkg/config"
	"github.com/soulmonk/cuppa-workers-authentication/pkg/db"
	"github.com/soulmonk/cuppa-workers-authentication/pkg/logger"
	"github.com/soulmonk/cuppa-workers-authentication/pkg/protocol/grpc"
	"github.com/soulmonk/cuppa-workers-authentication/pkg/protocol/rest"
	"log"
)

// RunServer runs gRPC server and HTTP gateway
func RunServer() error {
	// new line
	ctx := context.Background()

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	cfg := config.Load()

	// initialize logger
	if err := logger.Init(cfg.LogLevel, cfg.LogTimeFormat); err != nil {
		return fmt.Errorf("failed to initialize logger: %v", err)
	}
	dao := db.GetDao(ctx, cfg.PostgresqlConnectionString)

	defer func() {
		if err := dao.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	// todo rest temporary not needed
	// run HTTP gateway
	go func() {
		_ = rest.RunServer(ctx, cfg.HTTPPort)
	}()

	return grpc.RunServer(ctx, dao, cfg.GRPCPort)
}
