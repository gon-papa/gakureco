
-- +migrate Up
CREATE TABLE IF NOT EXISTS users (
  id serial PRIMARY KEY,
  name varchar(255) NOT NULL,
	email varchar(255) NOT NULL UNIQUE,
	password  varchar NOT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp,
  deleted_at timestamp
);

-- +migrate Down
DROP TABLE IF EXISTS users;