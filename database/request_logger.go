package database

import (
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"justQit/constants"
	"justQit/types"
	"log"
	"os"
	"strings"
)

var (
	reqCount       int = 0
	logFile            = "database/tmp.jsonl"
	reqLimit       int = 2
	LoggerQueue    chan *types.DBSchema
	DB             *sqlx.DB
	LogUpdateQueue chan *types.DBSchema
)

func backgroundLogAddQueue() {
	for req := range LoggerQueue {
		// Write to file
		reqJSON, err := json.Marshal(req)
		if err != nil {
			fmt.Println("Unable to log request " + req.JobID)
		}
		f, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		if _, err := f.Write(append(reqJSON, '\n')); err != nil {
			log.Fatal(err)
		}
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
		reqCount++
		// Increment counter, see if it's reached capacity
		if reqCount == reqLimit {
			// Extract from file, write to DB
			f, err := os.ReadFile(logFile)
			if err != nil {
				log.Fatal(err)
			}
			reqs := strings.Split(string(f), "\n")
			dbEntries := make([]types.DBSchema, 0, reqLimit) // len=0, cap=reqLimit
			for _, req := range reqs {
				var entry types.DBSchema
				if err := json.Unmarshal([]byte(req), &entry); err != nil {
					continue
				}
				dbEntries = append(dbEntries, entry)
			}
			ok := LogToDatabase(dbEntries)
			if ok == 0 {
				fmt.Println("Error logging to database")
				// TODO Create a DLQ file & store
			}
			os.Truncate(logFile, 0)
			reqCount = 0
		}
	}
}

func backgroundLogUpdateQueue() {}

func InitializeLogger(logAfterXreqs int) {
	var err error

	DB, err = sqlx.Open("sqlite3", "database/data.db")
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}

	_, err = DB.Exec(constants.LOG_DB_SCHEMA_SQL)
	if err != nil {
		log.Fatalf("failed to create table: %v", err)
	}

	log.Println("Database initialized and jobs table ready.")
	reqLimit = logAfterXreqs
	LoggerQueue = make(chan *types.DBSchema, constants.LOG_BUFFER_SIZE)
	go backgroundLogAddQueue()
	go backgroundLogUpdateQueue()
	fmt.Println("Initialized request Logger!")
}
