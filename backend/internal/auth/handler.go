package auth

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	"internal-ops-portal/internal/users"
)

type Handler struct {
	Users users.Repository
}

func NewHandler(usersRepo users.Repository) *Handler {
	return &Handler{Users: usersRepo}
}

func (h *Handler) Signup(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	hash, err := HashPassword(req.Password)
	if err != nil {
		http.Error(w, "failed to hash password", http.StatusInternalServerError)
		return
	}

	user := &users.User{
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: hash,
		Role:         "employee",
	}

	if err := h.Users.Create(context.Background(), user); err != nil {
		http.Error(w, "email already exists", http.StatusConflict)
		return
	}

	token, _ := GenerateJWT(user.ID, user.Role, os.Getenv("JWT_SECRET"))

	json.NewEncoder(w).Encode(map[string]any{
		"token": token,
		"user": map[string]string{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
			"role":  user.Role,
		},
	})
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	user, err := h.Users.GetByEmail(context.Background(), req.Email)
	if err != nil || user == nil {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}

	if !CheckPassword(req.Password, user.PasswordHash) {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
	}

	token, _ := GenerateJWT(user.ID, user.Role, os.Getenv("JWT_SECRET"))

	json.NewEncoder(w).Encode(map[string]any{
		"token": token,
		"user": map[string]string{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
			"role":  user.Role,
		},
	})
}
