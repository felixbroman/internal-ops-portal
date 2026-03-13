package requests

import "time"

type Request struct {
	ID           string
	Type         string
	Title        string
	Description  *string
	Status       string
	CreatedBy    string
	AssignedTo   *string
	DecisionBy   *string
	DecisionNote *string
	StartAt      *time.Time
	EndAt        *time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
