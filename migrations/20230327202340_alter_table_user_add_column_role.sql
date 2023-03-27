-- +goose Up
-- +goose StatementBegin
ALTER TABLE users
    ADD COLUMN role VARCHAR(8) NOT NULL DEFAULT 'user';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users
    DROP COLUMN role;
-- +goose StatementEnd
