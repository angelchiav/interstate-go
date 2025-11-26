package posts

import (
	"context"

	db "github.com/angelchiav/interstate-go/internal/database"
	"github.com/google/uuid"
)

type Service struct {
	database *db.Queries
}

func NewService(db *db.Queries) *Service {
	return &Service{database: db}
}

func (s *Service) CreatePost(ctx context.Context, body string, userID uuid.UUID) (db.Post, error) {
	post, err := s.database.CreatePost(ctx, db.CreatePostParams{
		Body:   body,
		UserID: userID,
	})
	if err != nil {
		return db.Post{}, err
	}

	return post, nil
}
