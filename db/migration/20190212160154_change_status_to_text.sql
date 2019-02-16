-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE public.elems (
    id SERIAL PRIMARY KEY,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    date    timestamp with time zone,
    priority integer,
    description text,
    status  text
);
-- +goose Down
DROP TABLE elems
-- SQL in this section is executed when the migration is rolled back.
