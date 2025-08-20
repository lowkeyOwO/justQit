package database

import (
	"fmt"
	"justQit/types"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sqlx.DB

const schema string = `
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

const query string = `
	INSERT INTO jobs (
		job_id, priority, payload, arrival_time, dispatch_time, 
		ack_time, ack_worker_id, status, log
	) VALUES (
		:job_id, :priority, :payload, :arrival_time, :dispatch_time,
		:ack_time, :ack_worker_id, :status, :log
	)
	`

// init runs automatically when the package is imported
func init() {
	var err error

	DB, err = sqlx.Open("sqlite3", "database/data.db")
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}

	_, err = DB.Exec(schema)
	if err != nil {
		log.Fatalf("failed to create table: %v", err)
	}

	log.Println("Database initialized and jobs table ready.")
}

func LogToDatabase(dbEntries []types.DBSchema) int {
	tx, err := DB.Beginx()
	if err != nil {
		fmt.Println("Transaction begin failed:", err)
		return 0
	}

	successCount := 0

	for _, entry := range dbEntries {
		_, err := tx.NamedExec(query, entry)
		if err != nil {
			fmt.Println("Insert failed:", err)
			_ = tx.Rollback()
			return successCount
		}
		successCount++
	}

	if err := tx.Commit(); err != nil {
		fmt.Println("Commit failed:", err)
		return successCount
	}

	return successCount
}

func UpdateInDatabase(dbEntries []types.DBSchema) {

}
