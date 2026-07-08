-- +goose Up
CREATE TABLE app_configuration (
    key VARCHAR(255) PRIMARY KEY,
    value VARCHAR(1024)
);

-- +goose Down
DROP TABLE app_configuration;
