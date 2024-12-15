-- name: CreateTask :one
INSERT INTO tasks (
        title,
        description,
        is_done,
        parent_list_id,
        parent_task_id
    )
VALUES (
        sqlc.arg('title'),
        sqlc.narg('description'),
        COALESCE(sqlc.arg('is_done'), false),
        sqlc.narg('parent_list_id'),
        sqlc.narg('parent_task_id')
    )
RETURNING *;

-- name: UpdateTask :one
UPDATE tasks
SET title = COALESCE(sqlc.narg('title'), title),
    description = sqlc.narg('description'),
    is_done = COALESCE(sqlc.narg('is_done'), is_done),
    parent_list_id = sqlc.narg('parent_list_id'),
    parent_task_id = sqlc.narg('parent_task_id')
WHERE id = sqlc.arg('id')
RETURNING *;

-- name: DeleteTask :exec
DELETE FROM tasks
WHERE id = sqlc.arg('id');

-- name: GetAllTasks :many
SELECT *
FROM tasks;

-- name: GetTaskByID :one
SELECT *
FROM tasks
WHERE id = sqlc.arg('id');

-- name: GetTasksByParentListID :many
SELECT *
FROM tasks
WHERE parent_list_id = sqlc.arg('id');

-- name: GetTasksByParentTaskID :many
SELECT *
FROM tasks
WHERE parent_task_id = sqlc.arg('id');