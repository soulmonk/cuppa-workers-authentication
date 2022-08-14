// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: user.query.sql

package user

import (
	"context"
	"time"
)

const activate = `-- name: Activate :exec
UPDATE "user" SET enabled=True WHERE id = $1
`

func (q *Queries) Activate(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, activate, id)
	return err
}

const create = `-- name: Create :one
INSERT INTO "user" (name, email, password, enabled, created_at, updated_at)
VALUES ($1, $2, $3, false, now(), now())
RETURNING id, enabled, created_at, updated_at
`

type CreateParams struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateRow struct {
	ID        int64     `json:"id"`
	Enabled   bool      `json:"enabled"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (q *Queries) Create(ctx context.Context, arg CreateParams) (CreateRow, error) {
	row := q.db.QueryRow(ctx, create, arg.Name, arg.Email, arg.Password)
	var i CreateRow
	err := row.Scan(
		&i.ID,
		&i.Enabled,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const delete = `-- name: Delete :exec
UPDATE "user" SET enabled=False WHERE id = $1
`

func (q *Queries) Delete(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, delete, id)
	return err
}

const findById = `-- name: FindById :one
SELECT id, name, email, password, enabled, created_at, updated_at, refresh_token FROM "user" where id = $1
`

func (q *Queries) FindById(ctx context.Context, id int64) (User, error) {
	row := q.db.QueryRow(ctx, findById, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.Enabled,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.RefreshToken,
	)
	return i, err
}

const findByName = `-- name: FindByName :one
SELECT id, name, email, password, enabled, created_at, updated_at, refresh_token FROM "user" where name = $1
`

func (q *Queries) FindByName(ctx context.Context, name string) (User, error) {
	row := q.db.QueryRow(ctx, findByName, name)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.Enabled,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.RefreshToken,
	)
	return i, err
}

const list = `-- name: List :many
SELECT id, name, email, enabled, created_at, updated_at FROM "user" ORDER BY updated_at DESC
`

type ListRow struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Enabled   bool      `json:"enabled"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (q *Queries) List(ctx context.Context) ([]ListRow, error) {
	rows, err := q.db.Query(ctx, list)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListRow
	for rows.Next() {
		var i ListRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Email,
			&i.Enabled,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
