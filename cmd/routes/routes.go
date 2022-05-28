package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"os"
)

func SetupRoutes(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Accept,Authorization,Content-Type,X-CSRF-TOKEN",
		ExposeHeaders:    "Link",
		AllowCredentials: true,
		MaxAge:           300,
	}))

	endpoint := os.Getenv("ENDPOINT")

	api := app.Group(endpoint)

	api.Get("/", GetHealthCheckData)
}
