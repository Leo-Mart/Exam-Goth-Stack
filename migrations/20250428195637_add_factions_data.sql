-- +goose Up
-- +goose StatementBegin
INSERT INTO faction (id, factionname) VALUES (1, 'Alliance');
INSERT INTO faction (id, factionname) VALUES (2, 'Horde');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM faction;
-- +goose StatementEnd
