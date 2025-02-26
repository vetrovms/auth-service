package services

import (
	"auth/internal/logger"
	"auth/internal/models"
	"auth/internal/request"
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// RegisterService Сервіс реєстрації.
type RegisterService struct {
	repo repositorer
}

// NewRegisterService Конструктор сервіса реєстрації.
func NewRegisterService(repo repositorer) RegisterService {
	return RegisterService{
		repo: repo,
	}
}

// Register Виконує реєстрацію користувача.
func (s *RegisterService) Register(ctx context.Context, r request.AuthRequest) error {
	bs, err := bcrypt.GenerateFromPassword([]byte(r.Password), bcrypt.DefaultCost)
	if err != nil {
		logger.Log().Info(err)
		return errors.New(SomethingWentWrongMsg)
	}
	user := models.User{
		Email:    r.Email,
		Password: string(bs),
	}
	err = s.repo.Save(ctx, user)
	if err != nil {
		logger.Log().Info(err)
		return errors.New(SomethingWentWrongMsg)
	}
	return nil
}
