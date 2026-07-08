-- +goose Up
CREATE TABLE heroes (
    id SERIAL PRIMARY KEY,
    beneficiary_id INTEGER NOT NULL REFERENCES beneficiaries(id) ON DELETE RESTRICT,
    name VARCHAR(50),
    phone VARCHAR(20),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ DEFAULT NULL
);

CREATE INDEX idx_heroes_beneficiary_id ON heroes (beneficiary_id);

CREATE UNIQUE INDEX uk_heroes_active_phone ON heroes (phone) WHERE deleted_at IS NULL;

-- +goose Down
DROP TABLE heroes;
