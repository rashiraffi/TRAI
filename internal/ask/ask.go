package ask

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"github.com/spf13/viper"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func Ask(ctx context.Context, query string) {
	// Access your API key as an environment variable (see "Set up your API key" above)
	client, err := genai.NewClient(ctx, option.WithAPIKey(viper.GetString("api_key")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// For text-only input, use the gemini-pro model
	model := client.GenerativeModel("gemini-pro")
	instructions := genai.Text("ensure that the response is formatted neatly for terminal display, as it will be printed directly to the terminal for readability.")
	prompt := genai.Text(query)
	iter := model.GenerateContentStream(ctx, instructions, prompt)

	for {
		resp, err := iter.Next()
		if err == iterator.Done {
			fmt.Print("\n")
			break
		}
		if err != nil {
			if strings.Contains(err.Error(), "Error 400") {
				fmt.Println("Error 400: Bad Request")
			}
			log.Fatal("Me", err)
		}
		printResponse(resp)
	}
}

func printResponse(resp *genai.GenerateContentResponse) {
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				fmt.Print(part)
			}
		}
	}
}
