package routes

import (
	"go-erp-nlm-mongo/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/users", handlers.CreateUser)
	app.Get("/users", handlers.GetUsers)

	app.Post("/inventory", handlers.CreateInventory)
	app.Get("/inventory", handlers.GetInventory)

	app.Post("/sales", handlers.CreateSale)
	app.Get("/sales", handlers.GetSales)

	app.Post("/openai", handlers.ProcessNaturalQueryOpenAI)
	app.Post("/ollamai", handlers.ProcessNaturalQueryOllama)
}
