package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var app = pkg.Application{}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	app.Config = config.Load()
	app.PgDao = pg.GetDao(&app.Config.Pg)

	defer func() {
		if err := app.PgDao.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	r := mux.NewRouter()

	controllers.Init(&app, r)

	addr := ":3000"
	log.Println("listen on", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatal(err)
	}
}
