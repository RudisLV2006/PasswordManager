package sql

import (
	"database/sql"
	"fmt"

	_ "github.com/glebarez/go-sqlite"
)

func Connection(dbFile string) {
	db, err := sql.Open("sqlite", dbFile) // Ensure the path to DB is correct
	if err != nil {
		fmt.Println(err)
		fmt.Println("error one tripped")
		return
	}

	defer db.Close()

	var sqliteVersion string
	err = db.QueryRow("select sqlite_version()").Scan(&sqliteVersion)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(sqliteVersion)
}
