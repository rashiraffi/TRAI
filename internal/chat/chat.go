package chat

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"github.com/spf13/viper"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func Chat(ctx context.Context) {
	// Access your API key as an environment variable (see "Set up your API key" above)
	client, err := genai.NewClient(ctx, option.WithAPIKey(viper.GetString("api_key")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// For text-only input, use the gemini-pro model
	model := client.GenerativeModel("gemini-pro")
	cs := model.StartChat()

	model.SafetySettings = []*genai.SafetySetting{
		{
			Category:  0,
			Threshold: 4,
		},
		{
			Category:  1,
			Threshold: 4,
		},
		{
			Category:  2,
			Threshold: 4,
		},
		{
			Category:  3,
			Threshold: 4,
		},
		{
			Category:  4,
			Threshold: 4,
		},
		{
			Category:  5,
			Threshold: 4,
		},
		{
			Category:  6,
			Threshold: 4,
		},
		{
			Category:  7,
			Threshold: 4,
		},
		{
			Category:  8,
			Threshold: 4,
		},
		{
			Category:  9,
			Threshold: 4,
		},
		{
			Category:  10,
			Threshold: 4,
		},
	}

	for {

		var query string
		fmt.Print("Ask: ")
		reader := bufio.NewReader(os.Stdin)
		query, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		if query == "exit\n" {
			break
		} else if query == "\n" {
			continue
		}

		iter := cs.SendMessageStream(ctx, genai.Text(query))
		fmt.Print("AI: ")

		for {
			resp, err := iter.Next()
			if err == iterator.Done {
				fmt.Print("\n")
				break
			}
			if err != nil {
				if strings.Contains(err.Error(), "Error 400") {
					fmt.Println("Error 400: Bad Request", err.Error())
					break
				}
				fmt.Println("Error: ", err)
				break
			}
			if resp != nil {
				printResponse(resp)
			}
		}

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
