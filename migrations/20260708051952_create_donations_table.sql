-- +goose Up
CREATE TYPE donation_status AS ENUM ('active', 'end');

CREATE TABLE donations (
    id SERIAL PRIMARY KEY,
    sponsor_id INTEGER NOT NULL REFERENCES sponsors(id) ON DELETE RESTRICT,
    quota INTEGER,
    remain INTEGER,
    take TIMESTAMPTZ,
    location VARCHAR(500) DEFAULT 'Pusat Studi Pancasila',
    maps VARCHAR(100) DEFAULT 'https://maps.app.goo.gl/eesnA6CN5fAQrGfP9',
    message VARCHAR(1024),
    media VARCHAR(1024),
    status donation_status NOT NULL DEFAULT 'end',
    reported BOOL NOT NULL DEFAULT FALSE,
    charity BOOL NOT NULL DEFAULT FALSE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_donations_active_feed
ON donations (created_at DESC)
WHERE status = 'active' AND remain > 0;

-- +goose Down
DROP TABLE donations;
DROP TYPE IF EXISTS donation_status;
