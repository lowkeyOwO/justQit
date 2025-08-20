package utils

import (
	"justQit/types"
	"time"
)

func CreateLogParcel(job_id string, priority int, payload string) *types.DBSchema {
	return &types.DBSchema{
		JobID:       job_id,
		Priority:    priority,
		Payload:     payload,
		ArrivalTime: time.Now(),
	}
}
