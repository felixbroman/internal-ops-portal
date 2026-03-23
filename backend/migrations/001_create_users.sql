CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE users (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  org_id UUID NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,

  name TEXT NOT NULL,
  email TEXT NOT NULL,
  password_hash TEXT NOT NULL,

  role TEXT NOT NULL CHECK (role IN ('employee', 'manager', 'admin')),
  manager_id UUID REFERENCES users(id),
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),

  UNIQUE (org_id, email)
);

CREATE INDEX idx_users_manager_id ON users(manager_id);
CREATE INDEX idx_users_role ON users(role);