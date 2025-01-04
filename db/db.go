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

	creteUsersTable := `
	CREATE TABLE IF NOT EXISTS USERS (
	ID SERIAL PRIMARY KEY,
	EMAIL TEXT NOT NULL UNIQUE,
	PASSWORD TEXT NOT NULL,
	SALT TEXT
	)
	`
	_, err := DB.Exec(creteUsersTable)
	if (err != nil) {
		log.Fatalf("Could not create the \"USERS\" table: " + err.Error())
	}

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		date_time TIMESTAMP NOT NULL,
		user_id INTEGER,
	CONSTRAINT FK_EVENTS_USERS FOREIGN KEY (user_id)
    REFERENCES USERS(ID)
	)
	`

	_, err = DB.Exec(createEventsTable)
	if err != nil {
		log.Fatalf("Could not create the \"EVENTS\" table: " + err.Error())
	}

	createRegistrationsTable := `
	CREATE TABLE IF NOT EXISTS registrations (
		id SERIAL PRIMARY KEY,
		user_id INTEGER,
		event_id INTEGER,
	CONSTRAINT FK_REGISTRATIONS_USERS FOREIGN KEY (user_id)
    REFERENCES USERS(ID),
	CONSTRAINT FK_REGISTRATIONS_EVENTS FOREIGN KEY (event_id)
    REFERENCES USERS(ID)
	)
	`

	_, err = DB.Exec(createRegistrationsTable)
	if err != nil {
		log.Fatalf("Could not create the \"EVENTS\" table: " + err.Error())
	}
}


