package utils

import (
	"fmt"
	"strconv"
	"strings"

	"justQit/constants"
	"github.com/manifoldco/promptui"
	"justQit/types"
)

func DispatcherInitPrompt() types.DispatcherConfig {
	var config types.DispatcherConfig

	// Prompt string
	promptString := func(label string) string {
		prompt := promptui.Prompt{Label: label}
		result, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed: %v\n", err)
		}
		return result
	}

	// Prompt int with validation
	promptInt := func(label string, defaultVal int, min int) int {
		prompt := promptui.Prompt{
			Label:   fmt.Sprintf("%s (default: %d)", label, defaultVal),
			Default: fmt.Sprintf("%d", defaultVal),
			Validate: func(input string) error {
				v, err := strconv.Atoi(input)
				if err != nil {
					return fmt.Errorf("invalid number")
				}
				if v < min {
					return fmt.Errorf("must be at least %d", min)
				}
				return nil
			},
		}
		result, _ := prompt.Run()
		value, _ := strconv.Atoi(result)
		return value
	}

	// Prompt int slice with custom separator and validation
	promptIntSlice := func(label, sep string) []int {
		for {
			prompt := promptui.Prompt{
				Label: fmt.Sprintf("%s (use '%s' as separator)", label, sep),
			}
			result, _ := prompt.Run()
			parts := strings.Split(result, sep)
			var values []int
			valid := true
			for _, p := range parts {
				p = strings.TrimSpace(p)
				if p == "" {
					continue
				}
				num, err := strconv.Atoi(p)
				if err != nil && num > 0 {
					fmt.Println("Invalid number:", p)
					valid = false
					break
				}
				values = append(values, num)
			}
			if valid && len(values) > 0 {
				return values
			}
			fmt.Println("Please enter a valid list of integers.")
		}
	}

	// Prompt fields
	config.MaxWorkers = promptInt("Max Workers", constants.MAX_WORKERS, 1)
	config.MaxDispatch = promptInt("Maximum requests dispatched", constants.MAX_DISPATCH_REQ, 1)
	config.QueueSize = promptIntSlice("Queue Sizes", ",")
	config.Priority = promptIntSlice("Priority Dispatch Ratio", ":")

	// Length validation
	if len(config.QueueSize) != len(config.Priority) {
		fmt.Println("Queue Sizes and Priority Ratios must be of the same length.")
		return DispatcherInitPrompt() 
	}

	// Optional custom field mapping
	index, _, err := (&promptui.Select{
		Label: "Enable Custom Fields?",
		Items: []string{"No", "Yes"},
	}).Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
	}

	if index == 1 {
		config.CustomFields = true
		config.FieldMap.JobId = promptString("Custom Field - Job ID")
		config.FieldMap.Priority = promptString("Custom Field - Priority")
	}

	return config
}
