-- +goose Up
create table public.auth_user
(
    id               bigserial
        primary key,
    password         varchar(128)             not null,
    last_login       timestamp with time zone,
    username         varchar(50)              not null
        unique,
    is_superuser     boolean                  not null,
    is_staff         boolean                  not null,
    is_active        boolean                  not null,
    is_test          boolean                  not null,
    date_joined      timestamp with time zone not null
);

create index auth_user_username_6300eb99_like
    on public.auth_user (username varchar_pattern_ops);

-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DROP TABLE public.auth_user;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
