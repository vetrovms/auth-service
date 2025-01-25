package services

import (
	"auth/internal/database/repository"
	"auth/internal/logger"
	"auth/internal/models"
	"auth/internal/request"
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// IRegisterService Інтерфейс сервіса реєстрації.
type IRegisterService interface {
	Register(ctx context.Context, r request.AuthRequest) error
}

// RegisterService Сервіс реєстрації.
type RegisterService struct {
	repo repository.IRepo
}

// NewRegisterService Конструктор сервіса реєстрації.
func NewRegisterService(repo repository.IRepo) RegisterService {
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
