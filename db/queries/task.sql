-- name: GetAllTasks :many
SELECT * FROM tasks;

-- name: GetTaskByParentListID :many
SELECT * FROM tasks WHERE parent_list_id = $1;

-- name: GetTaskByParentTaskID :many
SELECT * FROM tasks WHERE parent_task_id = $1;

-- name: GetTaskByID :one
SELECT * FROM tasks WHERE id = $1;

-- name: CreateTask :one
INSERT INTO tasks (title, description, status, parent_list_id) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: DeleteTask :exec
DELETE FROM tasks WHERE id = $1;