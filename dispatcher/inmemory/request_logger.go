package inmemory

import (
	"fmt"
	// "justQit/data"
	"justQit/types"
	"time"
)

var loggerQueue chan string
var reqCount int = 0

func requestLogger(job_id string, priority int, payload string) {
	req := types.DBSchema{
		JobID:       job_id,
		Priority:    priority,
		Payload:     payload,
		ArrivalTime: time.Now().String(),
	}
	loggerQueue <- req.JobID

	if len(loggerQueue) == cap(loggerQueue) {
		// Write from file to DB
	} else {
		request := <-loggerQueue
		// got a request → write to file, increment counter
		fmt.Println(request)

	}
}
