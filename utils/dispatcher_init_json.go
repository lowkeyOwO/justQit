package utils

import (
	"encoding/json"
	"fmt"
	"justQit/types"
	"os"
)

func DispatcherReadJSON(jsonpath string) types.DispatcherConfig {
	fileContent, err := os.ReadFile(jsonpath)
	if err != nil {
		fmt.Println("error:\t", err.Error())
	}
	var config types.DispatcherConfig
	json.Unmarshal(fileContent, &config)
	return config
}
