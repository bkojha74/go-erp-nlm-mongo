package handlers

import (
	"context"
	"go-erp-nlm-mongo/config"
	"go-erp-nlm-mongo/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func GetUsers(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := config.DB.Collection("users").Find(ctx, bson.M{})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	var users []models.User
	if err := cursor.All(ctx, &users); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(users)
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	// Parse JSON body into user struct
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input: " + err.Error(),
		})
	}

	// Basic validation
	if user.Name == "" || user.Email == "" || user.Role == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Name, Email, and Role are required",
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Insert into MongoDB
	_, err := config.DB.Collection("users").InsertOne(ctx, user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create user: " + err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully",
		"user":    user,
	})
}
