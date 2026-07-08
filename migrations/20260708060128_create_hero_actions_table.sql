-- +goose Up
CREATE TYPE hero_action_status AS ENUM ('done', 'ongoing');

CREATE TABLE hero_actions (
    id SERIAL PRIMARY KEY,
    hero_id INTEGER NOT NULL REFERENCES heroes(id),
    donation_id INTEGER NOT NULL REFERENCES donations(id) ON DELETE RESTRICT,
    status hero_action_status DEFAULT 'ongoing',
    code VARCHAR(50),
    quantity INTEGER NOT NULL DEFAULT 1 CHECK (quantity > 0),
    weight INTEGER NOT NULL DEFAULT 1 CHECK (weight > 0),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_hero_actions_hero_id ON hero_actions (hero_id);
CREATE INDEX idx_hero_actions_donation_id ON hero_actions (donation_id);

-- +goose Down
DROP TABLE hero_actions;
DROP TYPE IF EXISTS hero_action_status;
