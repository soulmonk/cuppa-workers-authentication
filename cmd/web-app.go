package main

import (
	"../config"
	"../pkg"
	"../pkg/db/pg"
	"../pkg/services"
	"../pkg/web/rest"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	app := pkg.Application{}
	conf := config.Load()
	pgDao := pg.GetDao(&conf.Pg)

	app.Services = services.Init(pgDao)
	app.Config = conf

	defer func() {
		if err := pgDao.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	rest.Init(&app)
}
