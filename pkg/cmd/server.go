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

type PG struct {
	Host     string
	Port     string
	User     string
	Password string
	Dbname   string
}

// Config is configuration for Server
type Config struct {
	// gRPC server start parameters section
	// gRPC is TCP port to listen by gRPC server
	GRPCPort string

	// HTTP/REST gateway start parameters section
	// HTTPPort is TCP port to listen by HTTP/REST gateway
	HTTPPort string

	Db PG

	// Log parameters section
	// LogLevel is global log level: Debug(-1), Info(0), Warn(1), Error(2), DPanic(3), Panic(4), Fatal(5)
	LogLevel int
	// LogTimeFormat is print time format for logger e.g. 2006-01-02T15:04:05Z07:00
	LogTimeFormat string
}

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
	db := pg.GetDao(&cfg.Pg)

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	v1API := v1.NewAuthenticationServiceServer(db.GetDb())

	// run HTTP gateway
	go func() {
		_ = rest.RunServer(ctx, cfg.HTTPPort)
	}()

	return grpc.RunServer(ctx, v1API, cfg.GRPCPort)
}
