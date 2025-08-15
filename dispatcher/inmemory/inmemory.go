package inmemory

import (
	"fmt"
	"io"
	"justQit/types"
	"net/http"
	"justQit/utils"
)

type InMemoryDispatcher struct {
	config    types.DispatcherConfig
	jobqueues []chan int
}

func (inmemory *InMemoryDispatcher) Initialize(config types.DispatcherConfig) {
	inmemory.config = config
	for i := range inmemory.jobqueues {
		queueSize := inmemory.config.QueueSize[i]
		inmemory.jobqueues[i] = make(chan int, queueSize)
	}
}

func (inmemory *InMemoryDispatcher) Enqueue(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// Read the body into bytes
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read body", http.StatusInternalServerError)
		return
	}
	job := utils.ExtractPayload(string(body))
	fmt.Println(job)
	w.Write([]byte("Got it!"))
}

func (inmemory *InMemoryDispatcher) JobASAP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Dispatch called")
}

func (inmemory *InMemoryDispatcher) Ack(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Ack called")
}
