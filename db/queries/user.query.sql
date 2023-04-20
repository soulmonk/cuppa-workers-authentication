-- name: Create :one
INSERT INTO "user" (name, email, password, enabled, role, created_at, updated_at)
VALUES ($1, $2, $3, false, role, now(), now())
RETURNING id, enabled, created_at, updated_at;

-- name: List :many
SELECT id, name, email, enabled, created_at, updated_at FROM "user" ORDER BY updated_at DESC;

-- name: FindById :one
SELECT * FROM "user" where id = $1;

-- name: FindByName :one
SELECT * FROM "user" where name = $1;

-- name: Delete :exec
UPDATE "user" SET enabled=False WHERE id = $1;

-- name: Activate :exec
UPDATE "user" SET enabled=True WHERE id = $1;

-- -- name AddRefreshToken :one
-- INSERT INTO "refresh_token" (name, email, password, enabled, role, created_at, updated_at)
-- VALUES ($1, $2, $3, false, role, now(), now())
-- RETURNING id, enabled, created_at, updated_at;
--
-- -- name: FindRefreshToken :one
-- SELECT id, user_id FROM "refresh_token" where token = $1;
--
-- -- name: RemoveRefreshToken :exec
-- DELETE FROM "refresh_token" where id = $1;
--
-- -- name: ClearUserToken :exec
-- DELETE FROM "refresh_token" where user_id = $1 and source = $2;
