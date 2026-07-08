-- +goose Up
CREATE TABLE heroes (
    id SERIAL PRIMARY KEY,
    beneficiary_id INTEGER NOT NULL REFERENCES beneficiaries(id) ON DELETE RESTRICT,
    name VARCHAR(50),
    phone VARCHAR(20) UNIQUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_heroes_beneficiary_id ON heroes (beneficiary_id);

-- +goose Down
DROP TABLE heroes;
