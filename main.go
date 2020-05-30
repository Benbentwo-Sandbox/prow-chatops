package main

import (
	actions "github.com/sethvargo/go-githubactions"
	"os"
)

func main() {
	fruit := actions.GetInput("fruit")
	if fruit == "" {
		actions.Fatalf("missing input 'fruit'")
	}
	githubToken := os.Getenv("GITHUB_TOKEN")

	actions.AddMask(fruit)
	actions.AddMask(githubToken)
}
