package repository

import (
	"../domain"
	"errors"
	"github.com/jmoiron/sqlx"
)

type UserVerificationDao struct {
	db        *sqlx.DB
	tableName string
}

func CreateUserVerificationDao(db *sqlx.DB) *UserVerificationDao {
	dao := UserVerificationDao{db, "user-verification"}
	return &dao
}

func (dao *UserVerificationDao) Update() error {
	return errors.New("not implemented")
}

func (dao *UserVerificationDao) Create() error {
	return errors.New("not implemented")
}

func (dao *UserVerificationDao) FindById(id string) (domain.UserVerification, error) {
	return domain.UserVerification{}, errors.New("not implemented")
}

func (dao *UserVerificationDao) Delete(id string) error {
	return errors.New("not implemented")
}
