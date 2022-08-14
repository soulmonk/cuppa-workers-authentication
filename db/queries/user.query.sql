-- name: Create :one
INSERT INTO "user" (name, email, password, enabled, created_at, updated_at)
VALUES ($1, $2, $3, false, now(), now())
RETURNING id, enabled, created_at, updated_at;

-- name: List :many
SELECT id, name, email, enabled, created_at, updated_at FROM "user" ORDER BY updated_at DESC;

-- name: FindById :one
SELECT * FROM "user" where id = $1;

-- name: FindByName :one
SELECT * FROM "user" where name = $1;

-- name: DeleteUser :exec
UPDATE "user" SET enabled=False WHERE id = $1;
