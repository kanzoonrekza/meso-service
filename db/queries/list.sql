-- name: GetAllLists :many
SELECT * FROM lists;

-- name: GetListByID :one
SELECT * FROM lists WHERE id = $1;

-- name: CreateList :one
INSERT INTO lists (title, description) VALUES ($1, $2) RETURNING *;

-- name: UpdateList :one
UPDATE lists SET title = COALESCE($1, title), description = COALESCE($2, description) WHERE id = $3 RETURNING *;

-- name: DeleteList :exec
DELETE FROM lists WHERE id = $1;