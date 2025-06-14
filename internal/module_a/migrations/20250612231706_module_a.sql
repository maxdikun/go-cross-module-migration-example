-- +goose Up
-- +goose StatementBegin
CREATE TABLE users(
    id INT PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
