-- +goose Up
CREATE TABLE report_keys (
    id SERIAL PRIMARY KEY,
    filename VARCHAR(1024),
    code VARCHAR(50) UNIQUE
);

-- +goose Down
DROP TABLE report_keys;
