-- +goose Up
-- +goose StatementBegin
CREATE TABLE chats(
    id INT PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE chats;
-- +goose StatementEnd
