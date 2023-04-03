-- +goose Up
-- +goose StatementBegin
ALTER TABLE queries
    ADD COLUMN field_descriptions TEXT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE queries
    DROP COLUMN field_descriptions;
-- +goose StatementEnd
