-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE public.elems (
    id integer NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    date    timestamp with time zone,
    priority integer,
    description text,
    status  BOOLEAN
);
-- +goose Down
DROP TABLE elems
-- SQL in this section is executed when the migration is rolled back.
