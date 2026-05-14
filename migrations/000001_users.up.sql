CREATE SCHEMA messanger;

CREATE TABLE messanger.users (
    id            UUID                        PRIMARY KEY,
    username      VARCHAR(16) UNIQUE NOT NULL CHECK(char_length(username) BETWEEN 1 AND 16),
    password_hash TEXT               NOT NULL,
    created_at    TIMESTAMPTZ        NOT NULL
);
