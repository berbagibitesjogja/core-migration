-- +goose Up
CREATE TABLE job_slots (
    id SERIAL PRIMARY KEY,
    job_schedule_id INTEGER NOT NULL REFERENCES job_schedules(id),
    code VARCHAR(255) UNIQUE,
    name VARCHAR(50),
    need INTEGER,
    place_and_time VARCHAR(500),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- +goose Down
DROP TABLE job_slots;
