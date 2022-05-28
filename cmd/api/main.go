package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go_healtcheck/cmd/routes"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Environment file coulnd't loaded. Error: %v", err)
	}

	port := os.Getenv("PORT")

	app := fiber.New()

	routes.SetupRoutes(app)

	log.Fatal(app.Listen(port))
}
