-- name: CreateTask :one
INSERT INTO tasks (
        title,
        description,
        is_done,
        due_date,
        priority,
        parent_list_id,
        parent_task_id
    )
VALUES (
        sqlc.arg('title'),
        sqlc.narg('description'),
        COALESCE(sqlc.arg('is_done'), false),
        sqlc.narg('due_date'),
        COALESCE(sqlc.narg('priority'), 0),
        sqlc.narg('parent_list_id'),
        sqlc.narg('parent_task_id')
    )
RETURNING *;

-- name: UpdateTask :one
UPDATE tasks
SET title = COALESCE(sqlc.narg('title'), title),
    description = sqlc.narg('description'),
    is_done = COALESCE(sqlc.narg('is_done'), is_done),
    due_date = sqlc.narg('due_date'),
    priority = COALESCE(sqlc.narg('priority'), priority),
    parent_list_id = sqlc.narg('parent_list_id'),
    parent_task_id = sqlc.narg('parent_task_id')
WHERE id = sqlc.arg('id')
RETURNING *;

-- name: DeleteTask :execrows
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

-- name: GetParentTaskData :one
SELECT id,
    parent_task_id
FROM tasks
WHERE id = sqlc.arg('id');