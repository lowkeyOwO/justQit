package constants

const MAX_WORKERS int = 1

const MAX_DISPATCH_REQ int = 10

const LOG_BUFFER_SIZE int = 100

const WORKER_ID_HEADER_KEY = "x-wrkr-id"

const LOG_DB_SCHEMA_SQL string = `
	CREATE TABLE IF NOT EXISTS jobs (
		job_id TEXT PRIMARY KEY,
		priority INTEGER,
		payload TEXT,
		arrival_time TIMESTAMPTZ,
		dispatch_time TIMESTAMPTZ,
		ack_time TIMESTAMPTZ,
		ack_worker_id TEXT,
		status TEXT,
		log TEXT
	);`

const LOG_DB_INSERT_SQL string = `
	INSERT INTO jobs (
		job_id, priority, payload, arrival_time, dispatch_time, 
		ack_time, worker_id, status, log
	) VALUES (
		:job_id, :priority, :payload, :arrival_time, :dispatch_time,
		:ack_time, :worker_id, :status, :log
	)
	`
