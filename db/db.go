package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func InitDB() *sql.DB {
	connStr := "host=localhost port=5432 user=salmanfarhat dbname=chocolate_db sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Connection failed:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Ping failed:", err)
	}

	log.Println("✅ Connected to the database")
	return db
}
