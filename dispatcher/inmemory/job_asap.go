package inmemory

import (
	"fmt"
	"justQit/constants"
	"net/http"
	"time"
)

func (inmemory *InMemoryDispatcher) JobASAP(w http.ResponseWriter, r *http.Request) {
	// Create a job request
	job := &jobReq{
		jobRequest: make(chan *JobResponse, 1),
	}
	// Send to request Channel
	jobDecider.reqChan <- job

	// Wait for a job_id from the request channel
	response := <-job.jobRequest
	if !response.Ok {
		http.Error(w, "no jobs available", http.StatusNotFound)
		return
	}

	// If we have a job - move from jobmap to dispatchmap, update details
	jobLog, _ := inmemory.jobMap.Get(response.JobID)
	jobLog.DispatchTime = time.Now()
	jobLog.Status = "Dispatched"

	fmt.Fprintf(w, "dispatched job: %s, with payload: %s", response.JobID, jobLog.Payload)
	
	jobLog.WorkerID = r.Header[http.CanonicalHeaderKey(constants.WORKER_ID_HEADER_KEY)][0]
	fmt.Println("worker id" + jobLog.WorkerID)
	inmemory.dispatchMap.Set(response.JobID, jobLog)

	// delete from jobmap
	inmemory.jobMap.Delete(jobLog.JobID)

}
