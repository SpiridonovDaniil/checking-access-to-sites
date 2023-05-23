-- +goose Up
-- +goose StatementBegin
CREATE TABLE access
(
    site varchar NOT NULL UNIQUE,
    response_time integer NOT NULL
);
CREATE TABLE statistic
(
    endpoint varchar NOT NULL UNIQUE,
    number_of_requests integer NOT NULL
);
INSERT INTO statistic (endpoint) VALUES ('getTime', 'getMinTime', 'getMaxTime')
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE access;
DROP TABLE statistic;
-- +goose StatementEnd
