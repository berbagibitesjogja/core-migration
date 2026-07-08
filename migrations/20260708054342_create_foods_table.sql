-- +goose Up
CREATE TYPE food_unit AS ENUM ('gr', 'ml');

CREATE TABLE foods (
    id SERIAL PRIMARY KEY,
    donation_id INTEGER NOT NULL REFERENCES donations(id) ON DELETE CASCADE,
    name VARCHAR(50),
    quantity INTEGER DEFAULT 1 CHECK (quantity > 0),
    weight INTEGER DEFAULT 1 CHECK (weight > 0),
    unit food_unit NOT NULL DEFAULT 'gr',
    notes VARCHAR(500),
    expired BOOL NOT NULL DEFAULT FALSE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_foods_donation_id ON foods (donation_id);

-- +goose Down
DROP TABLE foods;
DROP TYPE IF EXISTS food_unit;
