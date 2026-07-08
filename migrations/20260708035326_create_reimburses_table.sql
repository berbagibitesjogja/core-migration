-- +goose Up
CREATE TABLE reimburses (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id),
    method VARCHAR(255) NOT NULL,
    target VARCHAR(255) NOT NULL,
    amount BIGINT NOT NULL CHECK (amount > 0),
    is_done BOOL NOT NULL DEFAULT FALSE,
    file VARCHAR(1024) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- +goose Down
DROP TABLE reimburses;
