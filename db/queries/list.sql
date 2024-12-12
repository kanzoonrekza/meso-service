-- name: GetAllLists :many
SELECT * FROM lists;

-- name: CreateList :one
INSERT INTO lists (title, description) VALUES ($1, $2) RETURNING *;

-- name: DeleteList :exec
DELETE FROM lists WHERE id = $1;