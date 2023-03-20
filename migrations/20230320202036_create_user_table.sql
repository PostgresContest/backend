-- +goose Up
-- +goose StatementBegin
CREATE TABLE users
(
    id            serial PRIMARY KEY,
    login         varchar(64) unique,
    password_hash varchar(128)                           not null,
    first_name    varchar(64),
    last_name     varchar(64),

    registered_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at    timestamp with time zone DEFAULT now() NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
