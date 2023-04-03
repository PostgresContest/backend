-- +goose Up
-- +goose StatementBegin
CREATE TABLE queries
(
    id           SERIAL PRIMARY KEY,
    query_raw    TEXT,
    query_hash   VARCHAR(64),
    result_raw   TEXT,
    result_hash  VARCHAR(64),

    created_at   TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL
);
CREATE INDEX idx_queries_query_hash ON queries (query_hash);
CREATE INDEX idx_queries_response_hash ON queries (result_hash);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE queries
-- +goose StatementEnd
