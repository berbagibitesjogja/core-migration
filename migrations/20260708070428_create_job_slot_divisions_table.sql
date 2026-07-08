-- +goose Up
CREATE TABLE job_slot_divisions (
    id SERIAL PRIMARY KEY,
    job_slot_id INTEGER NOT NULL REFERENCES job_slots(id),
    division_id INTEGER NOT NULL REFERENCES divisions(id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT unique_job_slot_divisions UNIQUE (job_slot_id, division_id)
);

CREATE INDEX idx_job_slot_divisions_division_id ON job_slot_divisions (division_id);

-- +goose Down
DROP TABLE job_slot_divisions;
