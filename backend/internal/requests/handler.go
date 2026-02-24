package requests

import (
	"encoding/json"
	"net/http"
	"time"

	"internal-ops-portal/internal/auth"
)

type Handler struct {
	Repo Repository
}

func NewHandler(repo Repository) *Handler {
	return &Handler{Repo: repo}
}

// POST /api/requests
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	user, ok := auth.FromContext(r.Context())

	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	var body struct {
		Type        string     `json:"type"`
		Title       string     `json:"title"`
		Description *string    `json:"description"`
		StartAt     *time.Time `json:"startAt"`
		EndAt       *time.Time `json:"endAt"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	req := &Request{
		Type:        body.Type,
		Title:       body.Title,
		Description: body.Description,
		CreatedBy:   user.ID,
		StartAt:     body.StartAt,
		EndAt:       body.EndAt,
	}

	// overlap check
	if req.StartAt != nil && req.EndAt != nil {
		if req.Type == "equipment" {
			hasOverlap, err := h.Repo.HasOverlap(
				r.Context(),
				req.Type,
				*req.StartAt,
				*req.EndAt,
			)
			if err != nil {
				http.Error(w, "server error", http.StatusInternalServerError)
				return
			}
			if hasOverlap {
				http.Error(w, "resource already booked", http.StatusConflict)
				return
			}
		}
	}

	if err := h.Repo.Create(r.Context(), req); err != nil {
		http.Error(w, "failed to create request", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(req)
}

// GET /api/requests/mine
func (h *Handler) Mine(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	user, ok := auth.FromContext(r.Context())
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	reqs, err := h.Repo.GetByCreator(r.Context(), user.ID)
	if err != nil {
		http.Error(w, "failed to fetch requests", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(reqs)
}

// GET /api/requests
func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	reqs, err := h.Repo.GetAll(r.Context())
	if err != nil {
		http.Error(w, "failed to fetch requests", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(reqs)
}

// PATCH /api/requests/{id}
func (h *Handler) UpdateDecision(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	user, ok := auth.FromContext(r.Context())
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "missing request id", http.StatusBadRequest)
		return
	}

	var body struct {
		Status       string  `json:"status"`
		DecisionNote *string `json:"decisionNote"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if body.Status != "approved" && body.Status != "rejected" {
		http.Error(w, "invalid status", http.StatusBadRequest)
		return
	}

	if err := h.Repo.UpdateDecision(
		r.Context(),
		id,
		body.Status,
		user.ID,
		body.DecisionNote,
	); err != nil {
		http.Error(w, "failed to update request", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
