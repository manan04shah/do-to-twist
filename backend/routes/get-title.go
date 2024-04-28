package routes

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	openai "github.com/sashabaranov/go-openai"
)

func getQuirkyTitle(todoBody string, apiKey string) (string, error) {
	// Create a new OpenAI client
	client := openai.NewClient(apiKey)

	// Set up the request parameters
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "Generate a quirky title for my to-do list: " + todoBody,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return "", err
	}

	// Extract the suggested title from the response
	suggestedTitle := resp.Choices[0].Message.Content

	return suggestedTitle, nil
}

func GetTitle(c *fiber.Ctx) error {
	var noteBody string

	if err := c.BodyParser(&noteBody); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("OPENAI_API_KEY")

	quirkyTitle, err := getQuirkyTitle(noteBody, apiKey)
	if err != nil {
		fmt.Println("Error:", err)
		return c.Status(500).JSON("Error generating quirky title")
	}

	fmt.Println("Quirky Title:", quirkyTitle)

	return c.Status(200).JSON(quirkyTitle)
}
