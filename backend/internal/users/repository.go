package users

import "context"

type Repository interface {
	Create(ctx context.Context, user *User) error
	GetByID(ctx context.Context, email string) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
}
