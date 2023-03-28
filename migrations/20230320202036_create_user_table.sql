-- +goose Up
-- +goose StatementBegin
CREATE TABLE users
(
    id            SERIAL PRIMARY KEY,
    login         VARCHAR(64) UNIQUE,
    password_hash VARCHAR(128)                           NOT NULL,
    first_name    VARCHAR(64),
    last_name     VARCHAR(64),

    registered_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    updated_at    TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
