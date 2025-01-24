package services

import (
	"auth/internal/database/repository"
	"auth/internal/logger"
	"auth/internal/request"
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

// Приватні константи.
const (
	emailOccupiedMsg = "email вже зареєстровано"
	emailAbsentMsg   = "email не знайдено"
	wrongPasswordMsg = "невірний пароль"
	maxMsg           = " довжина має бути не більше%s символів"
	minMsg           = " довжина має бути не менше%s символів"
	emailFormatMsg   = " невірний формат email%s"
	requiredMsg      = " обов'язкове поле%s"
)

// Публічні константи
const (
	SomethingWentWrongMsg = "щось пішло не так, спробуйте пізніше"
)

type IValidationService interface {
	ValidateRegister(ctx context.Context, r request.AuthRequest) ([]string, error)
	ValidateLogin(ctx context.Context, r request.AuthRequest) ([]string, error)
}

type ValidationService struct {
	repo repository.IRepo
}

func NewValidationService(repo repository.IRepo) ValidationService {
	return ValidationService{
		repo: repo,
	}
}

func (v *ValidationService) ValidateRegister(ctx context.Context, r request.AuthRequest) ([]string, error) {
	exists, err := v.repo.UserExistsByEmail(ctx, r.Email)
	if err != nil {
		logger.Log().Info(err)
		return nil, errors.New(SomethingWentWrongMsg)
	}
	if exists {
		return []string{emailOccupiedMsg}, nil
	}
	msgs := validate(r)
	return msgs, nil
}

func (v *ValidationService) ValidateLogin(ctx context.Context, r request.AuthRequest) ([]string, error) {
	user, err := v.repo.GetUserByEmail(ctx, r.Email)
	if err != nil {
		logger.Log().Info(err)
		return nil, errors.New(SomethingWentWrongMsg)
	}
	if user.ID == 0 {
		return []string{emailAbsentMsg}, nil
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(r.Password))
	if err != nil {
		return []string{wrongPasswordMsg}, nil
	}
	msgs := validate(r)
	return msgs, nil
}

func validate(r interface{}) []string {
	var res []string
	errMap := errorsMap()
	validate := validator.New()
	errs := validate.Struct(r)

	if errs != nil {
		res := make([]string, 0)
		for _, err := range errs.(validator.ValidationErrors) {
			field := strings.ToLower(err.StructField())
			msg := strings.TrimSpace(field + fmt.Sprintf(errMap[err.Tag()], " "+err.Param()))
			res = append(res, msg)
		}
		return res
	}
	return res
}

func errorsMap() map[string]string {
	return map[string]string{
		"max":      maxMsg,
		"min":      minMsg,
		"email":    emailFormatMsg,
		"required": requiredMsg,
	}
}
