-- Users
CREATE INDEX idx_users_org_id ON users(org_id);

-- Requests
CREATE INDEX idx_requests_org_id ON requests(org_id);
CREATE INDEX idx_requests_created_by ON requests(created_by);
CREATE INDEX idx_requests_status ON requests(status);

--Invites
CREATE INDEX idx_invites_org_id ON org_invites(org_id);
CREATE INDEX idx_invites_email ON org_invites(email);