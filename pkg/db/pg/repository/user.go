package repository

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/soulmonk/cuppa-workers-authentication/pkg/db/pg/domain"
	"log"
)

type UserDao struct {
	//base *BaseDao
	db        *sqlx.DB
	tableName string
}

func CreateUserDao(db *sqlx.DB) *UserDao {
	dao := UserDao{db, "user"}
	return &dao
}

func (dao *UserDao) Update(model *domain.User) error {
	return errors.New("not implemented")
}

func (dao *UserDao) Create(model *domain.User) error {
	query := `INSERT INTO "` + dao.tableName + `" (name, email, password, enabled, created_at, updated_at) 
VALUES ($1, $2, $3, false, now(), now())
RETURNING id, enabled, created_at, updated_at`
	err := dao.db.
		QueryRow(query, model.Name, model.Email, model.Password).
		Scan(&model.ID, &model.Enabled, &model.CreatedAt, &model.UpdatedAt)

	if err != nil {
		log.Println("Error on create model")
		return err
	}

	return nil
}

func (dao *UserDao) List() (*domain.Users, error) {
	var res = domain.Users{}
	var err error

	// todo wtf
	var resp = func(err error) (*domain.Users, error) {
		return &res, err
	}

	rows, err := dao.db.Queryx(`SELECT id, name, email, enabled, created_at, updated_at FROM "` + dao.tableName + `" ORDER BY updated_at DESC`)

	if err != nil {
		log.Println("Error on executing query")
		return resp(err)
	}

	defer func() {
		if err := rows.Close(); err != nil {
			log.Println("Error corrupted while closing rows:", err.Error())
		}
	}()

	for rows.Next() {
		model := domain.User{}
		if err := rows.StructScan(&model); err != nil {
			log.Println("Error corrupted while scanning user:", err.Error())
			return &res, err
		}

		res.List = append(res.List, model)
	}
	if err := rows.Err(); err != nil {
		log.Println("Error on fetching rows:", err.Error())
		return &res, err
	}
	return &res, err
}

func (dao *UserDao) FindById(id string) (*domain.User, error) {
	query := `SELECT * FROM "` + dao.tableName + `" where id = $1`
	var model = domain.User{}

	if err := dao.db.Get(&model, query, id); err != nil {
		log.Println("Error on fetching note", err.Error())
		return &model, err
	}
	return &model, nil
}

func (dao *UserDao) Delete(id string) error {
	query := `UPDATE "` + dao.tableName + `" SET enabled=False WHERE id = $1`
	if _, err := dao.db.Exec(query, id); err != nil {
		log.Println("Error on deleting user", err.Error())
		return err
	}
	return nil
}

func (dao *UserDao) FindByName(username string) (*domain.User, error) {
	query := `SELECT * FROM "` + dao.tableName + `" where name = $1`
	var model = domain.User{}

	if err := dao.db.Get(&model, query, username); err != nil {
		log.Println("Error on fetching user", err.Error())
		return &model, err
	}
	return &model, nil
}
