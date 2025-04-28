-- +goose Up
-- +goose StatementBegin
CREATE TABLE character (
    id VARCHAR(255) PRIMARY KEY,
    characterprofile INT NOT NULL,
    keystoneprofile INT NOT NULL,
    gear INT NOT NULL,
    media INT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE character;
-- +goose StatementEnd
