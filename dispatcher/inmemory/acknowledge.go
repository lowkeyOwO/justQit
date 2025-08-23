package inmemory

import (
	"fmt"
	"io"
	"justQit/database"
	"justQit/utils"
	"net/http"
	"time"
)

/*
Read jobID, ack_type & log from req body, update in db, delete from dispatch map
*/
func (inmemory *InMemoryDispatcher) Ack(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read body", http.StatusInternalServerError)
		return
	}
	ackPayload := utils.ExtractAck(body)
	dispatchLog, Ok := inmemory.dispatchMap.Get(ackPayload.JobId)
	if !Ok {
		fmt.Fprintf(w, "job does not exist")
		return
	}

	if ackPayload.Ack {
		dispatchLog.Status = "Done"
	} else {
		dispatchLog.Status = "Failed"
	}
	dispatchLog.AckTime = time.Now()
	dispatchLog.Log = ackPayload.Log
	database.LogUpdateQueue <- dispatchLog
	inmemory.dispatchMap.Delete(ackPayload.JobId)
}
