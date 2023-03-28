-- +goose Up
-- +goose StatementBegin
CREATE TABLE queries
(
    id            SERIAL PRIMARY KEY,
    query_raw     TEXT,
    query_hash    VARCHAR(64),
    response_raw  TEXT,
    response_hash VARCHAR(64),
    execute_time  FLOAT,

    created_at    TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL
);
CREATE INDEX idx_queries_query_hash on queries(query_hash);
CREATE INDEX idx_queries_response_hash on queries(query_hash);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE queries
-- +goose StatementEnd
