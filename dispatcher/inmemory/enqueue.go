package inmemory

import (
	"io"
	"net/http"

	"justQit/database"
	"justQit/utils"
)

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
		if len(inmemory.jobQueues[priority]) == cap(inmemory.jobQueues[priority]) {
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte("Too many requests, please try later!"))
		} else {
			logParcel := utils.CreateLogParcel(job_id, priority, payload)
			inmemory.jobMap.Set(job_id, logParcel)
			inmemory.jobQueues[priority] <- job_id
			w.WriteHeader(http.StatusAccepted)
			w.Write([]byte("Job written with ID:\t" + job_id))
			if inmemory.config.LogToDatabase {
				database.LoggerQueue <- logParcel
			}
		}
	}
}
