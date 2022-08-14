package main

import (
	"context"
	"github.com/soulmonk/cuppa-workers-authentication/db/migration"
	"github.com/soulmonk/cuppa-workers-authentication/pkg/config"
	"github.com/soulmonk/cuppa-workers-authentication/pkg/db"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Starting migration")
	cfg := config.Load()

	ctx := context.Background()
	connection := db.InitConnection(ctx, cfg.PostgresqlConnectionString)

	defer func() {
		if err := connection.Close(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	if err := migration.Proceed(connection); err != nil {
		log.Fatal(err)
	}
	log.Println("Done")
}
