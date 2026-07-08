-- +goose Up
CREATE TABLE beneficiaries (
    id SERIAL PRIMARY KEY,
    affiliation_id INTEGER NOT NULL REFERENCES affiliations(id),
    name VARCHAR(50),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_beneficiaries_affiliation_id ON beneficiaries (affiliation_id);

-- +goose Down
DROP TABLE beneficiaries;
