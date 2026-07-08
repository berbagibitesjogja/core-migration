-- +goose Up
CREATE TYPE user_role AS ENUM ('super', 'core', 'staff', 'member');

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    division_id INTEGER NOT NULL REFERENCES divisions(id),
    kratos_id UUID NOT NULL UNIQUE,
    role user_role NOT NULL DEFAULT 'member',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- +goose Down
DROP TABLE users;
DROP TYPE IF EXISTS user_role;
