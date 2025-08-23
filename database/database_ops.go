package database

import (
	"fmt"
	"justQit/constants"
	"justQit/types"
)

func LogToDatabase(dbEntries []types.DBSchema) int {
	tx, err := DB.Beginx()
	if err != nil {
		fmt.Println("Transaction begin failed:", err)
		return 0
	}

	successCount := 0

	for _, entry := range dbEntries {
		_, err := tx.NamedExec(constants.LOG_DB_INSERT_SQL, entry)
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
