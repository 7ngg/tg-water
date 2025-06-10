-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS user(
    telegram_id INTEGER PRIMARY KEY,
    chat_id INTEGER NOT NULL UNIQUE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user;
-- +goose StatementEnd
