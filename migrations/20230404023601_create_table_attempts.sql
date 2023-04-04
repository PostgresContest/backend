-- +goose Up
-- +goose StatementBegin
CREATE TABLE attempts
(
    id         SERIAL PRIMARY KEY,
    user_id    INTEGER REFERENCES users (id),
    query_id   INTEGER REFERENCES queries (id),
    task_id    INTEGER REFERENCES tasks (id),
    accepted   BOOLEAN                  DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE attempts;
-- +goose StatementEnd
