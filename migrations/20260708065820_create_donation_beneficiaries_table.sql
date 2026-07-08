-- +goose Up
CREATE TABLE donation_beneficiaries (
    id SERIAL PRIMARY KEY,
    donation_id INTEGER NOT NULL REFERENCES donations(id) ON DELETE CASCADE,
    beneficiary_id INTEGER NOT NULL REFERENCES beneficiaries(id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT unique_donation_beneficiaries UNIQUE (donation_id, beneficiary_id)
);

CREATE INDEX idx_donation_beneficiaries_beneficiary_id ON donation_beneficiaries (beneficiary_id);

-- +goose Down
DROP TABLE donation_beneficiaries;
