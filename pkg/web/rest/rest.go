package rest

import (
	"../../../pkg"
	"./controllers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Init(app *pkg.Application) {
	r := mux.NewRouter()

	controllers.Init(app, r)

	addr := ":" + app.Config.Port
	log.Println("listen on", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatal(err)
	}
}
