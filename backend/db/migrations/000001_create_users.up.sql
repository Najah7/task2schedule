CREATE EXTENSION IF NOT EXISTS citext;

CREATE TABLE users (
    -- Q: Why 26 characters? 
    -- A: We use ULID (Universally Unique Lexicographically Sortable Identifier) for user IDs.
    id text PRIMARY KEY CHECK (id ~ '^[0-9ABCDEFGHJKMNPQRSTVWXYZ]{26}$'),
    first_name text NOT NULL,
    last_name text NOT NULL,
    email citext NOT NULL UNIQUE,
    password text NOT NULL,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now()
);
