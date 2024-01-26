package main

import (
	envconf "tipen-demo/config/env"
	"tipen-demo/handler"
	"tipen-demo/middleware"
	"tipen-demo/repository"
	"tipen-demo/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	env := envconf.ReadEnv()
	// dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable"
	repository := repository.NewRepository(repository.RepositoryParams{
		Url: env.Database.Url(),
	})
	handler := handler.NewHandler(handler.HandlerParams{
		Repository: repository,
	})
	middleware := middleware.NewMiddleware()
	route := route.NewRoute(route.RouteParams{
		App:        app,
		Handler:    handler,
		Middleware: middleware,
	})
	route.InitUser()

	app.Listen(":3000")
}
