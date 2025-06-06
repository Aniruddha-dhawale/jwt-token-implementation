package initializers

import (
	"database/sql"
	"log"
	_ "modernc.org/sqlite"
)

var DB *sql.DB

func ConnectDatabase() {
	var err error
	DB, err = sql.Open("sqlite", "test.db") // Use modernc.org/sqlite driver
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Drop the old users table if it exists
	_, err = DB.Exec(`DROP TABLE IF EXISTS users`)
		if err != nil {
		log.Fatal("Failed to drop old users table:", err)
	}

	// Create users table if it doesn't exist
	_, err = DB.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL ,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL
		)
	`)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
}
