-- +goose Up
-- +goose StatementBegin
CREATE TABLE messages(
    id INT PRIMARY KEY,
    chat INT REFERENCES chats(id) NOT NULL,
    data VARCHAR(1024) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE messages;
-- +goose StatementEnd
