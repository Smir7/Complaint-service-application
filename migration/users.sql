-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
	id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	user_uuid UUID NOT NULL UNIQUE,
    username VARCHAR(64) NOT NULL UNIQUE,
    password TEXT NOT NULL,
    role VARCHAR(8) DEFAULT 'USER'
);
CREATE INDEX IF NOT EXISTS idx_id ON users (id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd