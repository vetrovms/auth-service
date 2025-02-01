package middleware

import (
	"auth/internal/controllers"
	"auth/internal/logger"
	"auth/internal/models"
	"auth/internal/response"
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
)

// Config Конфігурація посередника перевірки авторства статті.
type Config struct {
	Filter  func(c *fiber.Ctx) bool
	Service controllers.Validator
}

func ClientNew(config Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		var r models.Client
		if err := c.BodyParser(&r); err != nil {
			logger.Log().Info(err)
			r := response.NewResponse(fiber.StatusUnprocessableEntity, []string{err.Error()}, nil)
			return c.Status(fiber.StatusUnprocessableEntity).JSON(r)
		}

		err := config.Service.ValidateClient(ctx, r)
		if err != nil {
			r := response.NewResponse(fiber.StatusUnprocessableEntity, []string{err.Error()}, nil)
			return c.Status(fiber.StatusUnprocessableEntity).JSON(r)
		}

		return c.Next()
	}
}
