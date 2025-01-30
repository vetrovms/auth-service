package services

import (
	"auth/internal/models"
	"context"
)

// repositorer Інтерфейс репозиторія.
type repositorer interface {
	UserExistsByEmail(ctx context.Context, email string) (bool, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	Save(ctx context.Context, user models.User) error
}
