package db

import (
	"database/sql"
	"log"
    _ "github.com/jackc/pgx/v5/stdlib"
)


var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("pgx", "postgres://postgres:elif@localhost:5432/gorestapi")

	if err != nil {
		log.Fatalf("Could not connect to the database: " + err.Error())
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables () {
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		date_time TIMESTAMP NOT NULL,
		user_id INTEGER
	)
	`

	_, err := DB.Exec(createEventsTable)
	if err != nil {
		log.Fatalf("Could not create the table: " + err.Error())
	}
}


