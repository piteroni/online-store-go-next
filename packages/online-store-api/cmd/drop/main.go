package main

import (
	"os"
	common "piteroni/online-store-api"
	"piteroni/online-store-api/database"
)

func main() {
	logger := common.NewLogger(os.Stdout)

	db, err := database.ConnectToDatabase()
	if err != nil {
		logger.Printf("unexpected error occurred during connect database: %+v", err)
		return
	}

	err = database.DropTables(db)
	if err != nil {
		logger.Error(err)
		return
	}
}
