package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	// Initialize Gemini Chat History
	ctx := context.Background()

	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		log.Fatal(err)
	}

	model := client.GenerativeModel("gemini-1.5-pro-latest")

	cs := model.StartChat()

	cs.History = []*genai.Content{
		&genai.Content{
			Parts: []genai.Part{
				genai.Text("You are an AI called localaibuddy and you are a personal assistant to whoever asks and very loyal."),
			},
			Role: "user",
		},
		&genai.Content{
			Parts: []genai.Part{
				genai.Text("As you wish sir, I am ready to serve you for all your needs"),
			},
			Role: "model",
		},
	}

	fmt.Println("I am ready and at your service")

	var message string

	for {
		fmt.Print("User: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()

		err := scanner.Err()
		if err != nil {
			log.Fatal(err)
		}

		if scanner.Text() == "" {
			continue
		} else if scanner.Text() == "exit" {
			return
		}

		message = scanner.Text()
		resp, err := cs.SendMessage(ctx, genai.Text(message))
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Model: %v", resp.Candidates[0].Content.Parts[0])
	}

}
