-- +goose Up
-- +goose StatementBegin
CREATE TYPE status AS ENUM ('to do', 'in progress', 'done');

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
    status status NOT NULL DEFAULT 'to do',
    parent_list_id BIGINT REFERENCES lists(id),
    parent_task_id BIGINT REFERENCES tasks(id),
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS tasks;

DROP TABLE IF EXISTS lists;

DROP TYPE IF EXISTS status;
-- +goose StatementEnd
