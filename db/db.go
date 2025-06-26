package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func InitDB() *sql.DB {
	connStr := "host=localhost port=5432 user=salmanfarhat dbname=chocolate_db sslmode=disable"
	db, err := sql.Open("Postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
