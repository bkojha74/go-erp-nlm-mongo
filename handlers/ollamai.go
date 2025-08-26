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
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

type ollamaAPIRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

type ollamaAPIResponse struct {
	Response string `json:"response"`
}

func ProcessNaturalQueryOllama(c *fiber.Ctx) error {
	fmt.Println("Received request for Ollama processing")
	var req models.OllamaRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	prompt := fmt.Sprintf(`You are an assistant for an ERP system. If the user asks about a person, extract their name and check if they exist in the database. Query: "%s" Respond only with the name if found, or say "not found".`, req.Message)

	fmt.Println("Constructed Prompt:", prompt)

	ollamaReq := ollamaAPIRequest{
		Model:  "mistral",
		Prompt: prompt,
		Stream: false,
	}

	body, _ := json.Marshal(ollamaReq)
	client := &http.Client{Timeout: 10 * time.Second}

	url := fmt.Sprintf("http://%s/api/generate", os.Getenv("OLLAMA_URL"))

	request, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))
	request.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(request)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Ollama request failed"})
	}
	defer resp.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	var ollamaResp ollamaAPIResponse
	if err := json.Unmarshal(bodyBytes, &ollamaResp); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to parse Ollama response"})
	}

	extractedName := strings.TrimSpace(ollamaResp.Response)

	if idx := strings.Index(extractedName, "("); idx != -1 {
		extractedName = strings.TrimSpace(extractedName[:idx])
	}

	// Query MongoDB for user
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user models.User

	filter := bson.M{"name": bson.M{"$regex": "^" + extractedName + "$", "$options": "i"}}
	err = config.DB.Collection("users").FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return c.JSON(models.OllamaResponse{Response: fmt.Sprintf("No user named %s found in the database.", extractedName)})
	}

	return c.JSON(models.OllamaResponse{Response: fmt.Sprintf("Yes, %s exists in our database and their email is %s.", user.Name, user.Email)})
}
