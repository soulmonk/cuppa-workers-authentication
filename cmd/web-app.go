package main

import (
	"../config"
	"../pkg"
	"../pkg/db/pg"
	"../pkg/protocol/rest"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	app := pkg.Application{}
	conf := config.Load()
	pgDao := pg.GetDao(&conf.Pg)

	app.Config = conf

	defer func() {
		if err := pgDao.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	rest.Init(&app)
}
