// backend/internal/requests/postgres_repository.go
package requests

import (
	"context"
	"database/sql"
	"time"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

// CREATE REQUEST
func (r *PostgresRepository) Create(ctx context.Context, req *Request) error {
	const q = `
		INSERT INTO requests (
			type,
			title,
			description,
			status,
			created_by,
			start_at,
			end_at
		)
		VALUES ($1, $2, $3, 'pending', $4, $5, $6)
		RETURNING
			id,
			status,
			created_at,
			updated_at
	`

	return r.db.QueryRowContext(
		ctx,
		q,
		req.Type,
		req.Title,
		req.Description,
		req.CreatedBy,
		req.StartAt,
		req.EndAt,
	).Scan(
		&req.ID,
		&req.Status,
		&req.CreatedAt,
		&req.UpdatedAt,
	)
}

func (r *PostgresRepository) GetByCreator(
	ctx context.Context,
	userID string,
) ([]Request, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT
			id, type, title, description, status,
			created_by, assigned_to,
			decision_by, decision_note,
			start_at, end_at,
			created_at, updated_at
		FROM requests
		WHERE created_by = $1
		ORDER BY created_at DESC
	`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []Request
	for rows.Next() {
		var r Request
		if err := rows.Scan(
			&r.ID,
			&r.Type,
			&r.Title,
			&r.Description,
			&r.Status,
			&r.CreatedBy,
			&r.AssignedTo,
			&r.DecisionBy,
			&r.DecisionNote,
			&r.StartAt,
			&r.EndAt,
			&r.CreatedAt,
			&r.UpdatedAt,
		); err != nil {
			return nil, err
		}
		out = append(out, r)
	}

	return out, nil
}

// GET ALL REQUESTS - ADMIN AND MANAGER ONLY
func (r *PostgresRepository) GetAll(ctx context.Context) ([]Request, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT
			id, type, title, description, status,
			created_by, assigned_to,
			decision_by, decision_note,
			start_at, end_at,
			created_at, updated_at
		FROM requests
		ORDER BY created_at DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []Request
	for rows.Next() {
		var r Request
		if err := rows.Scan(
			&r.ID,
			&r.Type,
			&r.Title,
			&r.Description,
			&r.Status,
			&r.CreatedBy,
			&r.AssignedTo,
			&r.DecisionBy,
			&r.DecisionNote,
			&r.StartAt,
			&r.EndAt,
			&r.CreatedAt,
			&r.UpdatedAt,
		); err != nil {
			return nil, err
		}
		out = append(out, r)
	}

	return out, nil
}

// UPDATE DECISION / STATUS
func (r *PostgresRepository) UpdateDecision(
	ctx context.Context,
	id string,
	status string,
	decisionBy string,
	note *string,
) error {
	_, err := r.db.ExecContext(ctx, `
		UPDATE requests
		SET
			status = $1,
			decision_by = $2,
			decision_note = $3,
			updated_at = now()
		WHERE id = $4
	`,
		status,
		decisionBy,
		note,
		id,
	)

	return err
}

// OVERLAP DETECTION
func (r *PostgresRepository) HasOverlap(
	ctx context.Context,
	reqType string,
	start time.Time,
	end time.Time,
) (bool, error) {
	const q = `
		SELECT 1
		FROM requests
		WHERE type = $1
		  AND status IN ('pending', 'approved')
		  AND start_at IS NOT NULL
		  AND end_at IS NOT NULL
		  AND start_at < $2
		  AND end_at > $3
		LIMIT 1
	`

	var exists int
	err := r.db.QueryRowContext(ctx, q, reqType, end, start).Scan(&exists)

	if err == sql.ErrNoRows {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	return true, nil
}
