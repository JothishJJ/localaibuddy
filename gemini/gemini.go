package gemini

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func GenerateModel() (*genai.GenerativeModel, context.Context, *genai.Client) {
	ctx := context.Background()

	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		log.Fatal(err)
	}

	model := client.GenerativeModel("gemini-1.5-pro-latest")

	return model, ctx, client
}

func CreateChat(model *genai.GenerativeModel) *genai.ChatSession {
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

	return cs
}

func SendMessage(message string, cs *genai.ChatSession, ctx context.Context) *genai.GenerateContentResponse {
	resp, err := cs.SendMessage(ctx, genai.Text(message))
	if err != nil {
		log.Fatal(err)
	}

	return resp
}
