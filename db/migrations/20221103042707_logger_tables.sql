-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
create table if not exists info (
    id bigserial not null primary key,
    date date,
    time time,
    id_tg int8,
    request text
);
create table if not exists errors (
    id bigserial not null primary key,
    date date,
    time time,
    id_tg int8,
    error text
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
drop table if exists info, errors;
-- +goose StatementEnd
