package handlers

import (
	"context"
	"go-erp-nlm-mongo/config"
	"go-erp-nlm-mongo/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateSale(c *fiber.Ctx) error {
	var sale models.Sale
	if err := c.BodyParser(&sale); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := config.DB.Collection("sales").InsertOne(ctx, sale)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{"message": "Sale recorded"})
}

func GetSales(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := config.DB.Collection("sales").Find(ctx, bson.M{}, options.Find())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	var sales []models.Sale
	if err := cursor.All(ctx, &sales); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(sales)
}
