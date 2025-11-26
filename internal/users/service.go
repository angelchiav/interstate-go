package users

import (
	"context"
	"fmt"

	db "github.com/angelchiav/interstate-go/internal/database"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	database *db.Queries
}

func NewService(db *db.Queries) *Service {
	return &Service{database: db}
}

func (s *Service) Register(ctx context.Context, username, password string) (db.User, error) {
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
