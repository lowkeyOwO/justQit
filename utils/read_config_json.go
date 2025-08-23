package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

// ReadJSON reads a JSON file and unmarshals it into any struct type T
func ReadJSON[T any](jsonPath string) T {
	var config T

	fileContent, err := os.ReadFile(jsonPath)
	if err != nil {
		fmt.Println("error reading file:\t", err.Error())
		os.Exit(1)
	}

	if err := json.Unmarshal(fileContent, &config); err != nil {
		fmt.Println("error unmarshaling JSON:\t", err.Error())
		os.Exit(1)
	}

	return config
}
