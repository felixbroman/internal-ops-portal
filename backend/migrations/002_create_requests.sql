CREATE TABLE requests (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  org_id UUID NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,

  type TEXT NOT NULL CHECK (type IN ('vacation', 'equipment', 'other')),
  title TEXT NOT NULL,
  description TEXT,

  status TEXT NOT NULL CHECK (status IN ('pending', 'approved', 'rejected')),

  created_by UUID NOT NULL REFERENCES users(id),
  assigned_to UUID REFERENCES users(id),
  decision_by UUID REFERENCES users(id),

  start_at DATE,
  end_at DATE,

  decision_note TEXT,

  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);
ADD CONSTRAINT requests_valid_time_window;
CHECK (
    start_at IS NULL
    OR end_at IS NULL
    OR end_at >= start_at
);

CREATE INDEX idx_requests_created_by ON requests(created_by);
CREATE INDEX idx_requests_assigned_to ON requests(assigned_to);