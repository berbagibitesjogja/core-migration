-- +goose Up
SELECT 'Note Begins';

-- Production deployed using podman version 5.8.2 on debian 13
-- Db is checked using DBeaver CE 26.1.0

-- +goose Down
SELECT 'Note Ends';
