package inmemory

import (
	"fmt"
	"io"
	"justQit/types"
	"justQit/utils"
	"net/http"
)

type InMemoryDispatcher struct {
	config    types.DispatcherConfig
	jobqueues []chan string
	jobmap    map[string]string
}

func (inmemory *InMemoryDispatcher) Initialize(config types.DispatcherConfig) {
	inmemory.config = config
	inmemory.jobqueues = make([]chan string, len(inmemory.config.QueueSize))

	for i := range inmemory.jobqueues {
		queueSize := inmemory.config.QueueSize[i]
		inmemory.jobqueues[i] = make(chan string, queueSize)
	}

	inmemory.jobmap = make(map[string]string)

	loggerQueue = make(chan string, config.LogAfterXRequests)
}

func (inmemory *InMemoryDispatcher) Enqueue(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// Read the body into bytes
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read body", http.StatusInternalServerError)
		return
	}
	job_id, priority, payload, err := utils.ExtractPayload(body, inmemory.config)
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		inmemory.jobmap[job_id] = payload
		if len(inmemory.jobqueues[priority]) == cap(inmemory.jobqueues[priority]) {
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte("Too many requests, please try later!"))
		} else {
			inmemory.jobqueues[priority] <- job_id
			w.WriteHeader(http.StatusAccepted)
			go requestLogger(job_id, priority, payload);
			w.Write([]byte("Job written with ID:\t" + job_id))
		}
	}
}

func (inmemory *InMemoryDispatcher) JobASAP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Dispatch called")
}

func (inmemory *InMemoryDispatcher) Ack(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Ack called")
}
