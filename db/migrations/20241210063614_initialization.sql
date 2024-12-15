-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS lists (
    id BIGSERIAL PRIMARY KEY,
    title TEXT NOT NULL CHECK (length(trim(title)) > 0),
    description TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS tasks (
    id BIGSERIAL PRIMARY KEY,
    title TEXT NOT NULL CHECK (length(trim(title)) > 0),
    description TEXT,
    is_done BOOLEAN NOT NULL DEFAULT false,
    parent_list_id BIGINT REFERENCES lists (id) ON DELETE CASCADE,
    parent_task_id BIGINT REFERENCES tasks (id) ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS tasks;

DROP TABLE IF EXISTS lists;
-- +goose StatementEnd