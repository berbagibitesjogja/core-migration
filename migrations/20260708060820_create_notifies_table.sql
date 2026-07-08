-- +goose Up
CREATE TABLE notifies (
    id SERIAL PRIMARY KEY,
    hero_id INTEGER NOT NULL UNIQUE REFERENCES heroes(id),
    notify_remains INTEGER DEFAULT 3,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- +goose Down
DROP TABLE notifies;
