// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: task.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createTask = `-- name: CreateTask :one
INSERT INTO tasks (
        title,
        description,
        is_done,
        parent_list_id,
        parent_task_id
    )
VALUES (
        $1,
        $2,
        COALESCE($3, false),
        $4,
        $5
    )
RETURNING id, title, description, is_done, parent_list_id, parent_task_id, created_at
`

type CreateTaskParams struct {
	Title        string      `json:"title"`
	Description  pgtype.Text `json:"description"`
	IsDone       interface{} `json:"is_done"`
	ParentListID pgtype.Int8 `json:"parent_list_id"`
	ParentTaskID pgtype.Int8 `json:"parent_task_id"`
}

func (q *Queries) CreateTask(ctx context.Context, arg CreateTaskParams) (Task, error) {
	row := q.db.QueryRow(ctx, createTask,
		arg.Title,
		arg.Description,
		arg.IsDone,
		arg.ParentListID,
		arg.ParentTaskID,
	)
	var i Task
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.IsDone,
		&i.ParentListID,
		&i.ParentTaskID,
		&i.CreatedAt,
	)
	return i, err
}

const deleteTask = `-- name: DeleteTask :exec
DELETE FROM tasks
WHERE id = $1
`

func (q *Queries) DeleteTask(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteTask, id)
	return err
}

const getAllTasks = `-- name: GetAllTasks :many
SELECT id, title, description, is_done, parent_list_id, parent_task_id, created_at
FROM tasks
`

func (q *Queries) GetAllTasks(ctx context.Context) ([]Task, error) {
	rows, err := q.db.Query(ctx, getAllTasks)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Task{}
	for rows.Next() {
		var i Task
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Description,
			&i.IsDone,
			&i.ParentListID,
			&i.ParentTaskID,
			&i.CreatedAt,
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

const getTaskByID = `-- name: GetTaskByID :one
SELECT id, title, description, is_done, parent_list_id, parent_task_id, created_at
FROM tasks
WHERE id = $1
`

func (q *Queries) GetTaskByID(ctx context.Context, id int64) (Task, error) {
	row := q.db.QueryRow(ctx, getTaskByID, id)
	var i Task
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.IsDone,
		&i.ParentListID,
		&i.ParentTaskID,
		&i.CreatedAt,
	)
	return i, err
}

const getTasksByParentListID = `-- name: GetTasksByParentListID :many
SELECT id, title, description, is_done, parent_list_id, parent_task_id, created_at
FROM tasks
WHERE parent_list_id = $1
`

func (q *Queries) GetTasksByParentListID(ctx context.Context, id pgtype.Int8) ([]Task, error) {
	rows, err := q.db.Query(ctx, getTasksByParentListID, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Task{}
	for rows.Next() {
		var i Task
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Description,
			&i.IsDone,
			&i.ParentListID,
			&i.ParentTaskID,
			&i.CreatedAt,
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

const getTasksByParentTaskID = `-- name: GetTasksByParentTaskID :many
SELECT id, title, description, is_done, parent_list_id, parent_task_id, created_at
FROM tasks
WHERE parent_task_id = $1
`

func (q *Queries) GetTasksByParentTaskID(ctx context.Context, id pgtype.Int8) ([]Task, error) {
	rows, err := q.db.Query(ctx, getTasksByParentTaskID, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Task{}
	for rows.Next() {
		var i Task
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Description,
			&i.IsDone,
			&i.ParentListID,
			&i.ParentTaskID,
			&i.CreatedAt,
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

const updateTask = `-- name: UpdateTask :one
UPDATE tasks
SET title = COALESCE($1, title),
    description = $2,
    is_done = COALESCE($3, is_done),
    parent_list_id = $4,
    parent_task_id = $5
WHERE id = $6
RETURNING id, title, description, is_done, parent_list_id, parent_task_id, created_at
`

type UpdateTaskParams struct {
	Title        pgtype.Text `json:"title"`
	Description  pgtype.Text `json:"description"`
	IsDone       pgtype.Bool `json:"is_done"`
	ParentListID pgtype.Int8 `json:"parent_list_id"`
	ParentTaskID pgtype.Int8 `json:"parent_task_id"`
	ID           int64       `json:"id"`
}

func (q *Queries) UpdateTask(ctx context.Context, arg UpdateTaskParams) (Task, error) {
	row := q.db.QueryRow(ctx, updateTask,
		arg.Title,
		arg.Description,
		arg.IsDone,
		arg.ParentListID,
		arg.ParentTaskID,
		arg.ID,
	)
	var i Task
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.IsDone,
		&i.ParentListID,
		&i.ParentTaskID,
		&i.CreatedAt,
	)
	return i, err
}
