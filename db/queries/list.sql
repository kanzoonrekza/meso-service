-- name: CreateList :one
INSERT INTO lists (title, description)
VALUES (sqlc.arg('title'), sqlc.narg('description'))
RETURNING *;

-- name: UpdateList :one
UPDATE lists
SET title = COALESCE(sqlc.narg('title'), title),
    description = sqlc.narg('description')
WHERE id = sqlc.arg('id')
RETURNING *;

-- name: DeleteList :exec
DELETE FROM lists
WHERE id = sqlc.arg('id');

-- name: GetAllLists :many
SELECT *
FROM lists;

-- name: GetListByID :one
SELECT *
FROM lists
WHERE id = sqlc.arg('id');