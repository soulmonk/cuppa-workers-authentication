package pg

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"

	"github.com/soulmonk/cuppa-workers-authentication/pkg/db/pg/repository"

	_ "github.com/lib/pq"
)

type Dao struct {
	UserDao             *repository.UserDao
	UserVerificationDao *repository.UserVerificationDao
	db                  *sqlx.DB
}

func InitConnection(connectionString string) *sqlx.DB {
	var err error
	//psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
	//	"password=%s dbname=%s sslmode=disable",
	//	config.Host, config.Port,
	//	config.User, config.Password, config.Dbname)

	fmt.Println(connectionString)

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
	dao.initModels()

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

	fmt.Println("Successfully connected!")
}

func (pg *Dao) initModels() {
	pg.UserDao = repository.CreateUserDao(pg.db)
	pg.UserVerificationDao = repository.CreateUserVerificationDao(pg.db)
}

// TODO not used circular because dependency
func (pg *Dao) Delete(from string, id string, modelName string) error {
	query := `DELETE FROM ` + from + ` WHERE id = $1`
	if _, err := pg.db.Exec(query, id); err != nil {
		log.Println("Error on deleting "+modelName, err.Error())
		return err
	}
	return nil
}

// TODO not used circular because dependency
func (pg *Dao) FindMyId(from string, id string, model *interface{}, modelName string) error {
	query := `SELECT * FROM "` + from + `" WHERE id = $1`
	if err := pg.db.Get(model, query, id); err != nil {
		log.Println("Error on fetching "+modelName, err.Error())
		return err
	}
	return nil
}
