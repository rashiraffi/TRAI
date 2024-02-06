package ask

import (
	"context"
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"github.com/spf13/viper"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

type AskParams struct {
	Query string
	IsCmd bool
	Run   bool
}

func Ask(ctx context.Context, params AskParams) {

	// Access your API key as an environment variable (see "Set up your API key" above)
	client, err := genai.NewClient(ctx, option.WithAPIKey(viper.GetString("api_key")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// For text-only input, use the gemini-pro model
	model := client.GenerativeModel("gemini-pro")
	var instructions []genai.Part
	var prompt genai.Text
	if params.IsCmd {
		prompt = genai.Text(fmt.Sprintf("user: %s", params.Query))
		instructions = append(instructions, cmdPromptParts...)
		instructions = append(instructions, prompt)
	} else {
		prompt = genai.Text(params.Query)
		instructions = []genai.Part{prompt}
	}

	iter := model.GenerateContentStream(ctx, instructions...)

	for {
		resp, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			if strings.Contains(err.Error(), "Error 400") {
				fmt.Println("Error 400: Bad Request")
			}
			log.Fatal("Me", err)
		}

		if params.IsCmd && params.Run && false { // TODO: Remove false when ready to run commands
			cmdStr := getResponse(resp)
			fmt.Println("Running command:", cmdStr)
			cmd := exec.Command("bash", "-c", cmdStr)
			err := cmd.Run()
			if err != nil {
				fmt.Println("Error running command:", err.Error())
			}
		} else {
			printResponse(resp)
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

func getResponse(resp *genai.GenerateContentResponse) string {
	res := ""
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				res += fmt.Sprintf("%v", part)
			}
		}
	}
	return res
}
