-- +goose Up
-- +goose StatementBegin
CREATE TABLE faction (
    id INT PRIMARY KEY,
    factionname VARCHAR(255) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE faction;
-- +goose StatementEnd
