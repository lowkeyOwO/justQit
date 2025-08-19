package types

type DBSchema struct {
	JobID        string `db:"job_id"`
	Priority     int    `db:"priority"`
	Payload      string `db:"payload"`
	ArrivalTime  string `db:"arrival_time"`
	DispatchTime string `db:"dispatch_time"`
	AckTime      string `db:"ack_time"`
	AckWorkerID  string `db:"ack_worker_id"`
	Status       string `db:"status"`
	Log          string `db:"log"`
}
