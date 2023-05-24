-- +goose Up
-- +goose StatementBegin
CREATE TABLE access
(
    site varchar NOT NULL UNIQUE,
    response_time bigint NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE access;
DROP TABLE statistic;
-- +goose StatementEnd
