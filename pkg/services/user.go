package services

import (
	"../db/pg/domain"
	"../db/pg/repository"
	"errors"
)

type UserService struct {
	dao *repository.UserDao
}

func CreateUserService(dao *repository.UserDao) *UserService {
	service := UserService{dao}
	return &service
}

func (service *UserService) Create(user *domain.User) error {
	return errors.New("not implemented")
}

func (service *UserService) List() (*domain.Users, error) {
	return service.dao.List()
}
