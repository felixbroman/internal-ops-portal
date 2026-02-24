package requests

import (
	"context"
	"time"
)

type Repository interface {
	Create(ctx context.Context, r *Request) error

	GetByCreator(ctx context.Context, userID string) ([]Request, error)
	GetAll(ctx context.Context) ([]Request, error)

	UpdateDecision(
		ctx context.Context,
		id string,
		status string,
		decisionBy string,
		note *string,
	) error

	HasOverlap(
		ctx context.Context,
		reqType string,
		start time.Time,
		end time.Time,
	) (bool, error)
}
