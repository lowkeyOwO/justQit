package types

import ("time")

type DBSchema struct {
	JobID        string `db:"job_id"`
	Priority     int    `db:"priority"`
	Payload      string `db:"payload"`
	ArrivalTime  time.Time `db:"arrival_time"`
	DispatchTime time.Time `db:"dispatch_time"`
	AckTime      time.Time `db:"ack_time"`
	AckWorkerID  string `db:"ack_worker_id"`
	Status       string `db:"status"`
	Log          string `db:"log"`
}
