package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func ConectionDatabase() *sql.DB {
	connection := "user=postgres dbname=golang-webapp password=admin host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connection)

	if err != nil {
		log.Fatal("Connection database error: ", err.Error())
	}

	return db
}
