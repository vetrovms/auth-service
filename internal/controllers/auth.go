package controllers

import (
	"auth/internal/logger"
	"auth/internal/models"
	"auth/internal/request"
	"auth/internal/response"
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
)

// validation Інтерфейс сервіса реєстрації.
type Validator interface {
	ValidateRegister(ctx context.Context, r request.AuthRequest) ([]string, error)
	ValidateLogin(ctx context.Context, r request.AuthRequest) ([]string, error)
	ValidateRetrospective(ctx context.Context, r request.RetrospectiveRequest) (bool, error)
	ValidateClient(ctx context.Context, m models.Client) error
}

// registrator Інтерфейс сервіса реєстрації.
type Registrator interface {
	Register(ctx context.Context, r request.AuthRequest) error
}

// login Інтерфейс сервіса логіна.
type Loginer interface {
	Login(ctx context.Context, r request.AuthRequest) (*string, error)
}

// AuthController контролер аутентифікації.
type AuthController struct {
	validation Validator
	register   Registrator
	login      Loginer
}

// ConfigAuthController Колбек для налаштування контролера.
type ConfigAuthController func(c *AuthController)

// WithValidationService Повертає колбек ConfigAuthController.
func WithValidationService(s Validator) func(c *AuthController) {
	return func(c *AuthController) {
		c.validation = s
	}
}

// WithRegisterService Повертає колбек ConfigAuthController.
func WithRegisterService(s Registrator) func(c *AuthController) {
	return func(c *AuthController) {
		c.register = s
	}
}

// WithLoginService Повертає колбек ConfigAuthController.
func WithLoginService(s Loginer) func(c *AuthController) {
	return func(c *AuthController) {
		c.login = s
	}
}

// NewAuthController Конструктор контролера.
func NewAuthController(configs ...ConfigAuthController) AuthController {
	c := AuthController{}
	for _, cnf := range configs {
		cnf(&c)
	}
	return c
}

// Register Обробник роута реєстрації нового користувача.
// Register godoc
//
//	@Summary		Реєстрація нового користувача
//	@Description	Реєстрація нового користувача
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param          request body request.AuthRequest true "login request (email, password)"
//	@Success		200		{object}	response.DocRegisterResponse200
//	@Failure		422		{object}	response.DocRegisterResponse422
//	@Failure		500		{object}	response.DocRegisterResponse500
//	@Router			/register [post]
func (a *AuthController) Register(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var r request.AuthRequest
	if err := c.BodyParser(&r); err != nil {
		logger.Log().Info(err)
		r := response.NewResponse(fiber.StatusUnprocessableEntity, []string{err.Error()}, nil)
		return c.Status(fiber.StatusUnprocessableEntity).JSON(r)
	}

	msgs, err := a.validation.ValidateRegister(ctx, r)
	if err != nil {
		r := response.NewResponse(fiber.StatusInternalServerError, []string{err.Error()}, nil)
		return c.Status(fiber.StatusInternalServerError).JSON(r)
	}
	if msgs != nil {
		r := response.NewResponse(fiber.StatusUnprocessableEntity, msgs, nil)
		return c.Status(fiber.StatusUnprocessableEntity).JSON(r)
	}

	err = a.register.Register(ctx, r)
	if err != nil {
		r := response.NewResponse(fiber.StatusInternalServerError, []string{err.Error()}, nil)
		return c.Status(fiber.StatusUnprocessableEntity).JSON(r)
	}

	res := response.NewResponse(fiber.StatusOK, []string{}, nil)
	return c.Status(fiber.StatusOK).JSON(res)
}

// Login Обробник роута логіна.
// Login godoc
//
//		@Summary		Логін
//		@Description	Логін
//		@Tags			users
//		@Accept			json
//		@Produce		json
//	    @Param          request body request.AuthRequest true "login request (email, password)"
//		@Success		200		{object}	response.DocLoginResponse200
//		@Failure		422		{object}	response.DocLoginResponse422
//		@Failure		500		{object}	response.DocLoginResponse500
//		@Router			/login [post]
func (a *AuthController) Login(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var r request.AuthRequest
	if err := c.BodyParser(&r); err != nil {
		logger.Log().Info(err)
		r := response.NewResponse(fiber.StatusUnprocessableEntity, []string{err.Error()}, nil)
		return c.Status(fiber.StatusUnprocessableEntity).JSON(r)
	}

	msgs, err := a.validation.ValidateLogin(ctx, r)
	if err != nil {
		r := response.NewResponse(fiber.StatusInternalServerError, []string{err.Error()}, nil)
		return c.Status(fiber.StatusInternalServerError).JSON(r)
	}
	if msgs != nil {
		r := response.NewResponse(fiber.StatusUnprocessableEntity, msgs, nil)
		return c.Status(fiber.StatusUnprocessableEntity).JSON(r)
	}

	jwt, err := a.login.Login(ctx, r)
	if err != nil {
		r := response.NewResponse(fiber.StatusInternalServerError, []string{err.Error()}, nil)
		return c.Status(fiber.StatusInternalServerError).JSON(r)
	}

	res := response.NewResponse(fiber.StatusOK, []string{}, fiber.Map{"jwt": *jwt})
	return c.Status(fiber.StatusOK).JSON(res)
}

func (a *AuthController) Retrospective(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var r request.RetrospectiveRequest
	if err := c.BodyParser(&r); err != nil {
		logger.Log().Info(err)
		r := response.NewResponse(fiber.StatusUnprocessableEntity, []string{err.Error()}, nil)
		return c.Status(fiber.StatusUnprocessableEntity).JSON(r)
	}

	ok, err := a.validation.ValidateRetrospective(ctx, r)
	if err != nil {
		res := response.NewResponse(fiber.StatusInternalServerError, []string{err.Error()}, nil)
		return c.Status(fiber.StatusInternalServerError).JSON(res)
	}

	res := response.NewResponse(fiber.StatusOK, []string{}, fiber.Map{"result": ok})
	return c.Status(fiber.StatusOK).JSON(res)
}
