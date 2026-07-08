-- +goose Up
CREATE TYPE sponsor_variant AS ENUM ('company', 'individual');

CREATE TABLE sponsors (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE,
    variant sponsor_variant DEFAULT 'individual',
    address VARCHAR(500),
    email VARCHAR(255),
    phone VARCHAR(20),
    logo VARCHAR(1024),
    hidden BOOL NOT NULL DEFAULT FALSE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- +goose Down
DROP TABLE sponsors;
DROP TYPE IF EXISTS sponsor_variant;
