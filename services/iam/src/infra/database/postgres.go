// Package database is responsible for connecting to the database and closing the connection.
// it uses postgresql as the database.

package database

import (
	"database/sql"
	"os"
)

func ConnectDB() *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_CONNECTION_STRING"))
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}

func CloseDB(db *sql.DB) {
	err := db.Close()
	if err != nil {
		panic(err)
	}
}
