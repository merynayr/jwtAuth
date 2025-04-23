-- +goose Up
-- +goose StatementBegin

CREATE TABLE users (
  user_id UUID PRIMARY KEY
);

CREATE TABLE jwt (
  id UUID PRIMARY KEY,
  user_id UUID NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
  token_hash TEXT NOT NULL,
  access_jti UUID NOT NULL,
  ip TEXT NOT NULL,
  user_agent TEXT,
  created_at TIMESTAMP NOT NULL,
  expires_at TIMESTAMP NOT NULL
);

CREATE INDEX idx_access_jti ON jwt (access_jti);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS jwt;
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
