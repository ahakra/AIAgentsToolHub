package main

import "fmt"
import "AIAgentsToolHub/internal/runner"

func main() {
	input := map[string]interface{}{
		"x": 5,
		"y": 8,
	}

	output, err := runner.CLIRunTool("./bin/add", input)
	if err != nil {
		panic(err)
	}

	fmt.Println("Result:", output)
}
