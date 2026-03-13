ALTER TABLE requests
ADD COLUMN start_at TIMESTAMPTZ,
ADD COLUMN end_at   TIMESTAMPTZ;

ALTER TABLE requests
ADD CONSTRAINT requests_valid_time_window
CHECK (
    start_at IS NULL
    OR end_at IS NULL
    OR end_at >= start_at
);

CREATE INDEX idx_requests_start_at ON requests(start_at);
CREATE INDEX idx_requests_end_at ON requests(end_at);