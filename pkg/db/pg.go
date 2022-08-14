package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/soulmonk/cuppa-workers-authentication/db/user"
)

type Dao struct {
	db          *sqlx.DB
	UserQuerier user.Querier
}

func InitConnection(connectionString string) *sqlx.DB {
	var err error

	db, err := sqlx.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}

func GetDao(connectionString string) *Dao {
	dao := Dao{}
	dao.initConnection(connectionString)

	return &dao
}

func (pg *Dao) Close() error {
	return pg.db.Close()
}

func (pg *Dao) GetDb() *sqlx.DB {
	return pg.db
}

func (pg *Dao) initConnection(connectionString string) {
	pg.db = InitConnection(connectionString)

	pg.UserQuerier = user.New(pg.db)
	fmt.Println("Successfully connected!")
}
