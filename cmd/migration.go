package main

import (
	"../config"
	"../pkg/db/pg"
	"../pkg/db/pg/migration"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Starting migration")
	cfg := config.Load()

	db := pg.InitConnection(&cfg.Pg)

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
