package main

import (
	"go-erp-nlm-mongo/config"
	"go-erp-nlm-mongo/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	config.ConnectDB()

	app := fiber.New()
	app.Use(cors.New())

	routes.SetupRoutes(app)

	app.Listen(":3000")
}
