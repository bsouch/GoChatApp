-- +goose Up
CREATE TABLE users (
    user_id UUID PRIMARY KEY NOT NULL,
    user_name VARCHAR(20) NOT NULL,
    password BYTEA NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE users;