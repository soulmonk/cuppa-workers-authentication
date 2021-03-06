package cmd

import (
	"context"
	"fmt"
	"github.com/soulmonk/cuppa-workers-authentication/pkg/logger"
	"log"
	// postgres driver
	_ "github.com/lib/pq"
	"github.com/soulmonk/cuppa-workers-authentication/pkg/config"
	"github.com/soulmonk/cuppa-workers-authentication/pkg/db/pg"
	"github.com/soulmonk/cuppa-workers-authentication/pkg/protocol/grpc"
	"github.com/soulmonk/cuppa-workers-authentication/pkg/protocol/rest"
	"github.com/soulmonk/cuppa-workers-authentication/pkg/service/v1"
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
	dao := pg.GetDao(&cfg.Pg)

	defer func() {
		if err := dao.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	v1API := v1.NewAuthenticationServiceServer(dao)

	// todo rest temporary not needed
	// run HTTP gateway
	go func() {
		_ = rest.RunServer(ctx, cfg.HTTPPort)
	}()

	return grpc.RunServer(ctx, v1API, cfg.GRPCPort)
}
