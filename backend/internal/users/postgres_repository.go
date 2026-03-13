package users

import (
	"context"
	"database/sql"
	"errors"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostpresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) Create(ctx context.Context, u *User) error {
	const query = `
		INSERT INTO users (name, email, password_hash, role, manager_id)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at
	`
	return r.db.QueryRowContext(
		ctx,
		query,
		u.Name,
		u.Email,
		u.PasswordHash,
		u.Role,
		u.ManagerID,
	).Scan(&u.ID, &u.CreatedAt)
}

func (r *PostgresRepository) GetByEmail(ctx context.Context, email string) (*User, error) {
	const query = `
		SELECT
			id,
			name,
			email,
			password_hash,
			role,
			manager_id,
			created_at
		FROM users
		WHERE email = $1
	`

	row := r.db.QueryRowContext(ctx, query, email)

	var u User
	err := row.Scan(
		&u.ID,
		&u.Name,
		&u.Email,
		&u.PasswordHash,
		&u.Role,
		&u.ManagerID,
		&u.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &u, nil
}

func (r *PostgresRepository) GetByID(ctx context.Context, id string) (*User, error) {
	const query = `
		SELECT id, name, email, password_hash, role, manager_id, created_at
		FROM users
		WHERE id = $1
	`

	var u User
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&u.ID,
		&u.Name,
		&u.Email,
		&u.PasswordHash,
		&u.ManagerID,
		&u.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &u, nil
}
