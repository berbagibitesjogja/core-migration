-- +goose Up
CREATE TABLE cancellations (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id),
    banned TIMESTAMPTZ NOT NULL DEFAULT (NOW() + INTERVAL '3 days'),
    tries INTEGER NOT NULL DEFAULT 1 CHECK (tries > 0),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_cancellations_user_banned ON cancellations (user_id) WHERE banned IS NOT NULL;

-- +goose Down
DROP TABLE cancellations;
