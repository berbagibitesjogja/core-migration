-- +goose Up
CREATE TABLE job_schedules (
    id SERIAL PRIMARY KEY,
    sponsor_id INTEGER NOT NULL REFERENCES sponsors(id),
    code VARCHAR(255) UNIQUE,
    receiver VARCHAR(255) NOT NULL,
    date TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_job_schedules_sponsor_id ON job_schedules (sponsor_id);

-- +goose Down
DROP TABLE job_schedules;
