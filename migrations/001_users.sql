-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    ID UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email TEXT NOT NULL UNIQUE,
    bio TEXT,
    hashed_password TEXT NOT NULL,
    CREATED_AT TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UPDATED_AT TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;

-- +goose StatementEnd