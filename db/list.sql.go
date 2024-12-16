// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: list.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createList = `-- name: CreateList :one
INSERT INTO lists (title, description)
VALUES ($1, $2)
RETURNING id, title, description, created_at
`

type CreateListParams struct {
	Title       string      `json:"title"`
	Description pgtype.Text `json:"description"`
}

func (q *Queries) CreateList(ctx context.Context, arg CreateListParams) (List, error) {
	row := q.db.QueryRow(ctx, createList, arg.Title, arg.Description)
	var i List
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.CreatedAt,
	)
	return i, err
}

const deleteList = `-- name: DeleteList :execrows
DELETE FROM lists
WHERE id = $1
`

func (q *Queries) DeleteList(ctx context.Context, id int64) (int64, error) {
	result, err := q.db.Exec(ctx, deleteList, id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}

const getAllLists = `-- name: GetAllLists :many
SELECT id, title, description, created_at
FROM lists
`

func (q *Queries) GetAllLists(ctx context.Context) ([]List, error) {
	rows, err := q.db.Query(ctx, getAllLists)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []List{}
	for rows.Next() {
		var i List
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Description,
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

const getListByID = `-- name: GetListByID :one
SELECT id, title, description, created_at
FROM lists
WHERE id = $1
`

func (q *Queries) GetListByID(ctx context.Context, id int64) (List, error) {
	row := q.db.QueryRow(ctx, getListByID, id)
	var i List
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.CreatedAt,
	)
	return i, err
}

const updateList = `-- name: UpdateList :one
UPDATE lists
SET title = COALESCE($1, title),
    description = $2
WHERE id = $3
RETURNING id, title, description, created_at
`

type UpdateListParams struct {
	Title       pgtype.Text `json:"title"`
	Description pgtype.Text `json:"description"`
	ID          int64       `json:"id"`
}

func (q *Queries) UpdateList(ctx context.Context, arg UpdateListParams) (List, error) {
	row := q.db.QueryRow(ctx, updateList, arg.Title, arg.Description, arg.ID)
	var i List
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.CreatedAt,
	)
	return i, err
}
