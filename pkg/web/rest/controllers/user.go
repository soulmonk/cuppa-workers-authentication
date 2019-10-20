package controllers

import (
	"../../../services"
	"../response"
	"github.com/gorilla/mux"
	"net/http"
)

type UserController struct {
	services *services.Services
}

func (ctrl *UserController) Init(r *mux.Router) {
	r.HandleFunc("/api/user", ctrl.list).Methods("GET")
}

func (ctrl *UserController) list(w http.ResponseWriter, r *http.Request) {
	users, err := ctrl.services.UserService.List()
	if err != nil {
		response.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// TODO format for web

	response.RespondWithJson(w, http.StatusOK, users)
}
