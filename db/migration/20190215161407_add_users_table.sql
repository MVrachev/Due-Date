-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE public.tasks (
    id SERIAL PRIMARY KEY,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    owner   text,
    due_date    timestamp with time zone,
    priority integer,
    description text,
    status  text
);

CREATE TABLE public.users (
    id SERIAL PRIMARY KEY,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name       text,
    password   text
);
-- +goose Down
DROP TABLE tasks
DROP TABLE users
-- SQL in this section is executed when the migration is rolled back.
