package controllers

import (
	"../../../services"
	"../response"
	"github.com/gorilla/mux"
)

type AuthController struct {
	services *services.Services
}

func (ctrl *AuthController) Init(r *mux.Router) {
	r.HandleFunc("/api/auth/login", response.Ni).Methods("GET")
	r.HandleFunc("/api/auth/logout", response.Ni).Methods("GET")
	r.HandleFunc("/api/auth/me", response.Ni).Methods("GET")
	r.HandleFunc("/api/auth/token", response.Ni).Methods("GET")
	r.HandleFunc("/api/auth/signup", response.Ni).Methods("POST")
}
