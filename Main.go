package main

import (
	"errors"
	"os"
	"server/api/data_access"
)

func main() {
	dbFile := "sql/PassMangerDB.db"

	if _, err := os.Stat(dbFile); errors.Is(err, os.ErrNotExist) {
		data_access.Create(dbFile)
	}

}
