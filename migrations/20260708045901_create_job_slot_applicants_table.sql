-- +goose Up
CREATE TABLE job_slot_applicants (
    id SERIAL PRIMARY KEY,
    job_slot_id INTEGER NOT NULL REFERENCES job_slots(id),
    user_id INTEGER NOT NULL REFERENCES users(id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT unique_slot_applicant UNIQUE (job_slot_id, user_id)
);

CREATE INDEX idx_job_slot_applicants_user_id ON job_slot_applicants (user_id);

-- +goose Down
DROP TABLE job_slot_applicants;
