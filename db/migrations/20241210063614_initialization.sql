-- +goose Up
-- +goose StatementBegin
CREATE TYPE status AS ENUM ('to do', 'in progress', 'done');

CREATE TABLE IF NOT EXISTS lists (
    id BIGSERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS tasks (
    id BIGSERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT,
    status status NOT NULL DEFAULT 'to do',
    list_id BIGINT NOT NULL REFERENCES lists(id),
    tast_id BIGSERIAL NOT NULL REFERENCES tasks(id),
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS tasks;

DROP TABLE IF EXISTS lists;

DROP TYPE IF EXISTS status;
-- +goose StatementEnd
