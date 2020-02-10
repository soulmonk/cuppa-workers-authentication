package rest

import (
	"../../../pkg"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Init(app *pkg.Application) {
	r := mux.NewRouter()

	r.HandleFunc("/api/status", status).Methods("GET")

	addr := ":" + app.Config.Port
	log.Println("listen on", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatal(err)
	}
}

type statusResponse struct {
	Status string `bson:"status" json:"status"`
}

func status(w http.ResponseWriter, r *http.Request) {
	var data = statusResponse{"ok"}
	response.RespondWithJson(w, http.StatusOK, data)
}
