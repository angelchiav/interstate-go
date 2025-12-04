package users

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/angelchiav/interstate-go/internal/database"
	db "github.com/angelchiav/interstate-go/internal/database"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	database *db.Queries
}

func NewService(db *db.Queries) *Service {
	return &Service{database: db}
}

func (s *Service) handlerRegister(ctx context.Context, username, password string) (db.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return db.User{}, fmt.Errorf("the user cannot be created")
	}
	user, err := s.database.CreateUser(ctx, db.CreateUserParams{
		Username:       username,
		HashedPassword: string(hash),
	})
	if err != nil {
		return db.User{}, fmt.Errorf("the user cannot be created")
	}
	return user, nil
}

func (s *Service) handlerChangePasswordByUsername(ctx context.Context, new_password, username string) error {
	user, err := s.database.GetUserByUsername(ctx, username)
	if err != nil {
		return err
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(new_password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	err = s.database.UpdatePasswordById(ctx, database.UpdatePasswordByIdParams{
		ID:             user.ID,
		HashedPassword: string(hashed),
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) handlerLogin(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var params parameters

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&params); err != nil {
		http.Error(w, "couldn't decode parameters", http.StatusBadRequest)
		return
	}

	user, err := s.database.GetUserByUsername(r.Context(), params.Username)
	if err != nil {
		http.Error(w, "username doesn't exists", http.StatusNotFound)
		return
	}
}
