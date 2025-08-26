package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"go-erp-nlm-mongo/config"
	"go-erp-nlm-mongo/models"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

type AIRequest struct {
	Message string `json:"message"`
}

type OpenAIRequest struct {
	Model    string              `json:"model"`
	Messages []map[string]string `json:"messages"`
}

type OpenAIResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func ProcessNaturalQueryOpenAI(c *fiber.Ctx) error {
	var req AIRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	// Step 1: Send message to OpenAI
	prompt := fmt.Sprintf(`You are an assistant for an ERP system. If the user asks about a person, extract their name and check if they exist in the database. Query: "%s" Respond only with the name if found, or say "not found".`, req.Message)

	openaiReq := OpenAIRequest{
		Model: "gpt-3.5-turbo",
		Messages: []map[string]string{
			{"role": "user", "content": prompt},
		},
	}

	body, _ := json.Marshal(openaiReq)
	client := &http.Client{Timeout: 10 * time.Second}
	apiKey := os.Getenv("OPENAI_API_KEY")

	fmt.Printf("Sending request to OpenAI having key[%+v] with prompt:%s", apiKey, prompt)

	request, _ := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(body))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+apiKey)

	resp, err := client.Do(request)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "OpenAI request failed"})
	}
	defer resp.Body.Close()

	var aiResp OpenAIResponse
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("OpenAI raw response:", string(bodyBytes))

	json.NewDecoder(resp.Body).Decode(&aiResp)

	if len(aiResp.Choices) == 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "OpenAI did not return any choices. Please check your prompt or API key.",
		})
	}

	extractedName := aiResp.Choices[0].Message.Content

	// Step 2: Query MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user models.User
	err = config.DB.Collection("users").FindOne(ctx, bson.M{"name": extractedName}).Decode(&user)
	if err != nil {
		return c.JSON(fiber.Map{"response": fmt.Sprintf("No user named %s found in the database.", extractedName)})
	}

	return c.JSON(fiber.Map{
		"response": fmt.Sprintf("Yes, %s exists in our database and their email is %s.", user.Name, user.Email),
	})
}
