package services

import (
	"../db/pg"
)

type Services struct {
	*UserService
}

func Init(dao *pg.Dao) *Services {
	services := Services{}
	services.UserService = CreateUserService(dao.UserDao)

	return &services
}
