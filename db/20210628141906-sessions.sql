-- +migrate Up
CREATE TABLE IF NOT EXISTS sessions (
  id serial NOT NULL NOT NULL UNIQUE PRIMARY KEY,
	email varchar(255) NOT NULL UNIQUE,
  uuid varchar(255) NOT NULL UNIQUE,
  user_id int NOT NULL UNIQUE REFERENCES users (id),
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp,
  deleted_at timestamp
);

-- +migrate Down
DROP TABLE IF EXISTS sessions;