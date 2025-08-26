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

func CreateInventory(c *fiber.Ctx) error {
	var item models.Inventory
	if err := c.BodyParser(&item); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := config.DB.Collection("inventory").InsertOne(ctx, item)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{"message": "Inventory item created"})
}

func GetInventory(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := config.DB.Collection("inventory").Find(ctx, bson.M{}, options.Find())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	var items []models.Inventory
	if err := cursor.All(ctx, &items); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(items)
}
