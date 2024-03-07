-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd


ALTER TABLE product_items ADD COLUMN prod_sec_name text NOT NULL DEFAULT '';

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

ALTER TABLE product_items DROP COLUMN prod_sec_name;