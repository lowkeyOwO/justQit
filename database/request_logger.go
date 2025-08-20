package database

import (
	"encoding/json"
	"fmt"
	"justQit/constants"
	"justQit/types"
	"log"
	"os"
	"strings"
)

var (
	reqCount    int = 0
	logFile         = "database/tmp.jsonl"
	ReqLimit    int = 2
	LoggerQueue chan *types.DBSchema
)

func InitializeLogger() {
	LoggerQueue = make(chan *types.DBSchema, constants.LOG_BUFFER_SIZE)
	go func() {
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
			if reqCount == ReqLimit {
				// Extract from file, write to DB
				f, err := os.ReadFile(logFile)
				if err != nil {
					log.Fatal(err)
				}
				reqs := strings.Split(string(f), "\n")
				dbEntries := make([]types.DBSchema, 0, ReqLimit) // len=0, cap=reqLimit
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
					// TODO Create a temporary file & store
				}
				os.Truncate(logFile, 0)
				reqCount = 0
			}
		}
	}()
	fmt.Println("Initialized request Logger!")
}
