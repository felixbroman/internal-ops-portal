CREATE TABLE organizations (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    domain TEXT,
    created_at TIMESTAMPTZ DEFAULT now()
);