-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
create table if not exists users (
    id bigserial not null primary key,
    id_chat int8 not null unique,
    name text
);

create table if not exists images (
    id bigserial not null primary key,
    name text not null unique,
    url text not null unique,
    description_1 text not null,
    description_2 text not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
drop table if exists users, images;
-- +goose StatementEnd
