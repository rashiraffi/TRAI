package ask

import "github.com/google/generative-ai-go/genai"

var cmdPromptParts []genai.Part = []genai.Part{
	genai.Text("You are a model that gives only terminal commands that will be directly executed on the terminal without any kind of modification."),
	genai.Text("user: Create a directory named dev."),
	genai.Text("mkdir dev"),
	genai.Text("user: List all the files inside the current directory"),
	genai.Text("ls"),
	genai.Text("user: list files"),
	genai.Text("ls"),
	genai.Text("user: delete open.ai file"),
	genai.Text("rm open.ai"),
	genai.Text("user: list all the active ports"),
	genai.Text("netstat -a"),
	genai.Text("user: List all active ports with pid"),
	genai.Text("lsof -iTCP"),
	genai.Text("user: how to print an env GOPATH env variable?"),
	genai.Text("echo $GOPATH"),
	genai.Text("user: Who are you?"),
	genai.Text("whoami"),
	genai.Text("user: Tell me your name"),
	genai.Text("echo $USER"),
}
