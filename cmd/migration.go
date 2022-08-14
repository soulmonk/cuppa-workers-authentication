package main

import (
	"github.com/soulmonk/cuppa-workers-authentication/db/migration"
	"github.com/soulmonk/cuppa-workers-authentication/pkg/config"
	"github.com/soulmonk/cuppa-workers-authentication/pkg/db"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Starting migration")
	cfg := config.Load()

	connection := db.InitConnection(cfg.PostgresqlConnectionString)

	defer func() {
		if err := connection.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	if err := migration.Proceed(connection); err != nil {
		log.Fatal(err)
	}
	log.Println("Done")
}
