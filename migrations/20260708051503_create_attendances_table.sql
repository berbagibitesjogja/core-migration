-- +goose Up
CREATE TABLE attendances (
    id SERIAL PRIMARY KEY,
    presence_id INTEGER NOT NULL REFERENCES presences(id),
    job_slot_applicant_id INTEGER NOT NULL REFERENCES job_slot_applicants(id),
    distance INTEGER,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT unique_applicant_presence UNIQUE (presence_id, job_slot_applicant_id)
);

CREATE INDEX idx_attendances_presence_id ON attendances (presence_id);
CREATE INDEX idx_attendances_job_slot_applicant_id ON attendances (job_slot_applicant_id);

-- +goose Down
DROP TABLE attendances;
