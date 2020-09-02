package router

import (
	"github.com/FranciscoMendes10866/api-go/guards"
	"github.com/FranciscoMendes10866/api-go/handler"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
)

func SetupRoutes(app *fiber.App) {
	app.Use(middleware.Logger())
	app.Use(middleware.Compress())
	api := app.Group("/api/v1")
	auth := api.Group("/auth")
	auth.Get("/", guards.AuthRequired(), handler.GetUsers)
	auth.Post("/create", handler.CreateUser)
	auth.Post("/login", handler.LoginUser)
}
