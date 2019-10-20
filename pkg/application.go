package pkg

import (
	"../config"
	"../pkg/services"
)

type Application struct {
	Config   *config.Config
	Services *services.Services
}
