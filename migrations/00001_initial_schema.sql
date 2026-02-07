CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  name text NOT NULL,
  email text NOT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp
);
