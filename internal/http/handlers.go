package http

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/angelchiav/interstate-go/internal/posts"
	"github.com/angelchiav/interstate-go/internal/users"
	"github.com/google/uuid"
)

type Handler struct {
	users *users.Service
	posts *posts.Service
}

func NewHandler(usersSVC *users.Service) *Handler {
	return &Handler{users: usersSVC}
}

func (h *Handler) handlerRegisterUser(w http.ResponseWriter, r *http.Request) {
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
		return
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

func (h *Handler) handlerChangePassword(w http.ResponseWriter, r *http.Request) {
	type Request struct {
		Username    string `json:"username"`
		NewPassword string `json:"new_password"`
	}

	var req Request

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	if err := h.users.ChangePasswordByUsername(r.Context(), req.NewPassword, req.Username); err != nil {
		http.Error(w, "could not change the password", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "password changed"})
}

func (h *Handler) handlerCreatePost(w http.ResponseWriter, r *http.Request) {
	type Request struct {
		UserID uuid.UUID `json:"user_id"`
		Body   string    `json:"body"`
	}

	var req Request

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		http.Error(w, "invalid JSON body", http.StatusBadRequest)
		return
	}

	post, err := h.posts.CreatePost(r.Context(), req.Body, req.UserID)
	if err != nil {
		http.Error(w, "Post cannot be created", http.StatusBadRequest)
		return
	}

	type Response struct {
		ID        string    `json:"id"`
		UserID    string    `json:"user_id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Body      string    `json:"body"`
	}

	out := Response{
		ID:        post.ID.String(),
		UserID:    post.UserID.String(),
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
		Body:      post.Body,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(out)
}
