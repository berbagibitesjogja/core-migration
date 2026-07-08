-- +goose Up
CREATE TABLE divisions (
    id SERIAL PRIMARY KEY,
    name VARCHAR(24) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- +goose Down
DROP TABLE divisions;
