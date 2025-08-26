package main

import (
	"go-erp-nlm-mongo/config"
	"go-erp-nlm-mongo/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	config.ConnectDB()

	app := fiber.New()

	routes.SetupRoutes(app)

	app.Listen(":3000")
}
