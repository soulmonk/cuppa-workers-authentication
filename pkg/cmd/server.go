package cmd

import (
	"context"
	"flag"
	"fmt"
	"github.com/jmoiron/sqlx"

	//"github.com/jmoiron/sqlx" TODO
	"log"

	// postgres driver
	_ "github.com/lib/pq"

	"../../config"
	"../db/pg"

	"../protocol/grpc"
	"../service/v1"
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
	db := pg.GetDao(&cfg.Pg)

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	v1API := v1.NewAuthenticationServiceServer(db)

	// run HTTP gateway
	go func() {
		_ = rest.RunServer(ctx, cfg.HTTPPort)
	}()

	return grpc.RunServer(ctx, v1API, cfg.GRPCPort)
}
