package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Init() {
	var err error
	connStr := "postgres://username:password@localhost/dbname?sslmode=disable"
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Unable to connect to database:", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("Database ping failed:", err)
	}
}
