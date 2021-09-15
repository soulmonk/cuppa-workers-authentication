package main

import (
	"github.com/soulmonk/cuppa-workers-authentication/pkg/config"
	"github.com/soulmonk/cuppa-workers-authentication/pkg/db/pg"
	"github.com/soulmonk/cuppa-workers-authentication/pkg/db/pg/migration"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Starting migration")
	cfg := config.Load()

	db := pg.InitConnection(cfg.PostgresqlConnectionString)

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	if err := migration.Proceed(db); err != nil {
		log.Fatal(err)
	}
	log.Println("Done")
}
