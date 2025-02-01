package main

import (
	"auth/internal/config"
	"auth/internal/controllers"
	"auth/internal/database/connection"
	"auth/internal/database/repository"
	"auth/internal/logger"
	"auth/internal/middleware"
	"auth/internal/services"

	_ "auth/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

//	@title			Authorization service.
//	@version		1.0
//	@description	Authorization service.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		127.0.0.1:8001
//	@BasePath	/

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	db := connection.Db()
	repo := repository.NewRepo(db)

	vs := services.NewValidationService(&repo)
	rs := services.NewRegisterService(&repo)
	ls := services.NewLoginService(&repo)

	controller := controllers.NewAuthController(
		controllers.WithLoginService(&ls),
		controllers.WithRegisterService(&rs),
		controllers.WithValidationService(&vs),
	)

	app := fiber.New()
	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Use(middleware.ClientNew(middleware.Config{Service: &vs}))
	app.Post("/register", controller.Register)
	app.Post("/login", controller.Login)
	app.Post("/retrospective", controller.Retrospective)

	logger.Log().Fatal(app.Listen(":" + config.NewEnv().WebPort))
}
