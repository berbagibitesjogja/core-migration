-- +goose Up
CREATE TYPE affiliation_variant AS ENUM ('student', 'society', 'foundation');

CREATE TABLE affiliations (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50),
    variant affiliation_variant NOT NULL DEFAULT 'student',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ DEFAULT NULL
);

-- +goose Down
DROP TABLE affiliations;
DROP TYPE IF EXISTS affiliation_variant;
