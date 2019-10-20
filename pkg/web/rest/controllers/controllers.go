package controllers

import (
	"../../../../pkg"
	"github.com/gorilla/mux"
)

func Init(app *pkg.Application, r *mux.Router) {

	// TODO mux middleware

	// TODO is authenticated

	(&AuthController{app.Services}).Init(r)
	(&UserController{app.Services}).Init(r)
}
