-- +goose Up
-- +goose StatementBegin
CREATE TABLE tasks
(
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(64),
    description TEXT,
    query_id    INTEGER,

    created_at  TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    FOREIGN KEY (query_id) REFERENCES queries (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE tasks;
-- +goose StatementEnd
