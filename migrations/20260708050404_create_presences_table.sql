-- +goose Up
CREATE TYPE presences_status AS ENUM ('active', 'end');

CREATE TABLE presences (
    id SERIAL PRIMARY KEY,
    job_slot_id INTEGER REFERENCES job_slots(id),
    title VARCHAR(50),
    description VARCHAR(500),
    latitude NUMERIC(9, 6) NOT NULL,
    longitude NUMERIC(10, 6) NOT NULL,
    max_distance INTEGER,
    code VARCHAR(255),
    status presences_status DEFAULT 'active',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_presences_job_slot_id ON presences (job_slot_id);

-- +goose Down
DROP TABLE presences;
DROP TYPE IF EXISTS presences_status;
