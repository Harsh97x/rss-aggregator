-- +goose up

CREATE TABLE users (
id UUID PRIMARY KEY
create_at TIMESTAMP NOT NULL
update_at TIMESTAMP NOT NULL
name TEXT NOT NULL
);
-- +goose down
