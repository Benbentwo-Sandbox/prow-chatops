package main

import (
	actions "github.com/sethvargo/go-githubactions"
)

func main() {
	fruit := actions.GetInput("fruit")
	if fruit == "" {
		actions.Fatalf("missing input 'fruit'")
	}
	actions.AddMask(fruit)
}
