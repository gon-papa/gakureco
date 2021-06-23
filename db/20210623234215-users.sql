
-- +migrate Up
CREATE TABLE IF NOT EXISTS users (
  id int NOT NULL UNIQUE PRIMARY KEY,
  name varchar NOT NULL,
	email varchar NOT NULL UNIQUE,
	password  varchar NOT NULL,
  created_at timestamp,
  updated_at timestamp
);

-- +migrate Down
DROP TABLE IF EXISTS users;