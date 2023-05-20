// Package database is responsible for connecting to the database and closing the connection.
// it uses postgresql as the database.

package database

import (
	"database/sql"
	"iam/pkg/apperrors"
	"os"
)

func ConnectDB() *sql.DB {
	dbConnectionString := os.Getenv("DATABASE_CONNECTION_STRING")

	if dbConnectionString == "" {
		panic(apperrors.DatabaseConnectionStringNotSet)
	}

	db, err := sql.Open("postgres", dbConnectionString)

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
