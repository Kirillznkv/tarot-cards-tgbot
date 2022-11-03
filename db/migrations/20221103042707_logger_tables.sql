-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
create table if not exists info (
    id bigserial not null primary key,
    id_chat int8,
    name text,
    date date,
    time time,
    request text
);
create table if not exists errors (
    id bigserial not null primary key,
    id_chat int8,
    date date,
    time time,
    error text
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
drop table if exists info, errors;
-- +goose StatementEnd
