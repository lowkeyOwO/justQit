package inmemory

type JobResponse struct {
	JobID string
	Ok    bool
}
type jobReq struct {
	jobRequest chan *JobResponse
}

type nextJobDecider struct {
	reqChan            chan *jobReq
	tempPriorityHolder []int
	resetMeter         bool
}

var jobDecider *nextJobDecider = nil

// Use round-robin to choose a queue & jobs
func (inmemory *InMemoryDispatcher) jobRequestHandler() {
	// Initialize jobDecider if not already done
	if jobDecider == nil {
		jobDecider = &nextJobDecider{
			reqChan:            make(chan *jobReq),
			tempPriorityHolder: make([]int, len(inmemory.config.Priority)),
			resetMeter:         false,
		}
		copy(jobDecider.tempPriorityHolder, inmemory.config.Priority)
	}

	// Use round-robin to choose a queue & jobs

	for request := range jobDecider.reqChan {
		var jobResponse *JobResponse = nil
		jobDecider.resetMeter = false
	found:
		for idx, val := range jobDecider.tempPriorityHolder {
			if val > 0 {
				// Check the queue
				select {
				case job_id := <-inmemory.jobQueues[idx]:
					jobResponse = &JobResponse{
						JobID: job_id,
						Ok:    true,
					}
					// break the loop if a job is found
					break found
				default:
					jobDecider.tempPriorityHolder[idx]--
					jobDecider.resetMeter = jobDecider.resetMeter || (jobDecider.tempPriorityHolder[idx] > 0)
				}
			}
		}
		if jobResponse == nil {
			jobResponse = &JobResponse{
				JobID: "",
				Ok:    false,
			}
		}

		// Send response back to the requesting body
		request.jobRequest <- jobResponse

		if !jobDecider.resetMeter {
			// Reset the queue
			copy(jobDecider.tempPriorityHolder, inmemory.config.Priority)
			jobDecider.resetMeter = false
		}

	}
}

/*
for each request in reqChan
 - go through each queue, find the first queue that has a job & return job id
 - reduce each element in queue by 1
 - if no job given till end of queue, return not ok
 - reset queue after every element becomes 0 (bool, keep or till end if > 0, and when we reach end, check if the bool is true
 if not, reset)
*/
