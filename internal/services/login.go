package services

import (
	"auth/internal/config"
	"auth/internal/database/repository"
	"auth/internal/logger"
	"auth/internal/request"
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// tokenDuration Час життя токена авторизації.
const tokenDuration = 12 * time.Hour

// ILoginService Інтерфейс сервіса логіна.
type ILoginService interface {
	Login(ctx context.Context, r request.AuthRequest) (*string, error)
}

// LoginService Сервіс логіна.
type LoginService struct {
	repo repository.IRepo
}

// NewLoginService Конструктор сервіса логіна.
func NewLoginService(repo repository.IRepo) LoginService {
	return LoginService{
		repo: repo,
	}
}

// Login Повертає JWT токен в разі успішного логіна.
func (l *LoginService) Login(ctx context.Context, r request.AuthRequest) (*string, error) {
	user, err := l.repo.GetUserByEmail(ctx, r.Email)
	if err != nil {
		logger.Log().Info(err)
		return nil, errors.New(SomethingWentWrongMsg)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Duration(tokenDuration)).Unix(),
	})

	tokenString, err := token.SignedString([]byte(config.NewEnv().SecretKey))
	if err != nil {
		logger.Log().Info(err)
		return nil, errors.New(SomethingWentWrongMsg)
	}

	return &tokenString, nil
}
