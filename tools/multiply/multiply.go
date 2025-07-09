package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Input struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func main() {
	var input Input
	if err := json.NewDecoder(os.Stdin).Decode(&input); err != nil {
		fmt.Fprintf(os.Stderr, "decode error: %v\n", err)
		os.Exit(1)
	}
	mult := input.X * input.Y
	json.NewEncoder(os.Stdout).Encode(map[string]int{"multiply": mult})
}
