package main

import "server/api/sql"

func main() {
	dbFile := "sql/PassMangerDB.db"
	sql.Check(dbFile)

	sql.Connection(dbFile)

}
