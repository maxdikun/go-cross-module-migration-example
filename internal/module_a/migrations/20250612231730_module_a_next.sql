-- +goose Up
-- +goose StatementBegin
CREATE TABLE friends(
    first INT REFERENCES users(id) NOT NULL,
    second INT REFERENCES users(id) NOT NULL,
    UNIQUE(first, second)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE friends;
-- +goose StatementEnd
