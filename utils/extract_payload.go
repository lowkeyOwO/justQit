package utils

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/google/uuid"
	"justQit/types"
)

func ExtractPayload(reqBody []byte, config types.DispatcherConfig) (string, int, string, error) {
	var payload map[string]any
	err := json.Unmarshal(reqBody, &payload)
	if err != nil {
		return "", 0, "", err
	}

	var jobID string
	var priority int
	var payloadMap map[string]any

	if config.CustomFields {
		// --- Job ID ---
		rawJobID, ok := payload[config.FieldMap.JobId]
		if !ok || rawJobID == nil {
			return "", 0, "", errors.New("missing job ID field")
		}
		s, ok := rawJobID.(string)
		if !ok || s == "" {
			return "", 0, "", errors.New("invalid job ID: must be a non-empty string")
		}
		jobID = s

		// --- Priority ---
		rawPriority, ok := payload[config.FieldMap.Priority]
		if !ok || rawPriority == nil {
			return "", 0, "", errors.New("missing priority field")
		}
		switch v := rawPriority.(type) {
		case int:
			priority = v
		case float64:
			priority = int(v)
			if float64(priority) != v {
				return "", 0, "", errors.New("invalid priority: not an integer value")
			}
		case string:
			p, err := strconv.Atoi(v)
			if err != nil {
				return "", 0, "", errors.New("invalid priority: must be an integer string")
			}
			priority = p
		default:
			return "", 0, "", errors.New("invalid priority type")
		}

		// Retain the entire body as payload
		payloadMap = payload
	} else {
		// Not using custom fields
		jobID = uuid.New().String()
		priority = config.Priority[len(config.Priority)-1]
		payloadMap = payload
	}

	payloadBytes, err := json.Marshal(payloadMap)
	if err != nil {
		return "", 0, "", err
	}

	return jobID, priority, string(payloadBytes), nil
}
