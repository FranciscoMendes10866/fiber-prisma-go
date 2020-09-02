package main

import (
	"github.com/FranciscoMendes10866/api-go/router"
	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	"github.com/gofiber/helmet"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())
	app.Use(helmet.New())

	router.SetupRoutes(app)

	app.Listen(3333)
}
