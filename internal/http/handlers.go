package http

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/angelchiav/interstate-go/internal/users"
)

type Handler struct {
	users *users.Service
}

func NewHandler(usersSVC *users.Service) *Handler {
	return &Handler{users: usersSVC}
}

func (h *Handler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	type Request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var req Request

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	user, err := h.users.Register(r.Context(), req.Username, req.Password)
	if err != nil {
		http.Error(w, "could not create user", http.StatusBadRequest)
	}

	type Response struct {
		ID        string    `json:"id"`
		Username  string    `json:"username"`
		CreatedAt time.Time `json:"created_at"`
	}

	out := Response{
		ID:        user.ID.String(),
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(out)
}
