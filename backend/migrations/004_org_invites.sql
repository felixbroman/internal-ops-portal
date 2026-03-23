CREATE TABLE org_invites (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    org_id UUID NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,

    email TEXT NOT NULL,

    token TEXT NOT NULL UNIQUE,
    invited_by UUID REFERENCES users(id),

    expires_at TIMESTAMPTZ,
    accepted_at TIMESTAMPTZ,

    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);