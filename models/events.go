package models

import (
	"time"

	"eliferden.com/restapi/db"
)

type Event struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"time" binding:"required"`
	UserID      int64      `json:"user_id"`
}

var events []Event = []Event{}

func (e *Event) Save() error {
    query := `
    INSERT INTO events(name, description, location, date_time, user_id) 
    VALUES($1,$2,$3,$4,$5)
    RETURNING id`
 
    var id int64
    err := db.DB.QueryRow(query, e.Name, e.Description, e.Location, e.DateTime, e.UserID).Scan(&id)
    if err != nil {
        return err
    }
    e.ID = id
    return err
}

func GetAllEvents() ([]Event, error) {
	query := `SELECT * FROM events;`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID,	&event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)

	}
	return events, nil
}

func GetEventByID (id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = $1"
	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}
	return &event, nil

}

func (e *Event) UpdateEvent () error{
	query := `
	UPDATE events
	SET name = $1, description = $2, location = $3, date_time = $4
	WHERE id = $5
	`
	preparedStatement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer preparedStatement.Close()

	_ , err = preparedStatement.Exec(e.Name, e.Description, e.Location, e.DateTime, e.ID)
	return err
}

func (e Event) DeleteEvent (id int64) error {
	query := "DELETE FROM events WHERE id = $1"
	preparedStatement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer preparedStatement.Close()

	_ , err = preparedStatement.Exec(e.ID)
	return err
}
func (e Event) Register(userId int64) error{
	query := `
    INSERT INTO registrations(event_id, user_id) 
    VALUES($1,$2)
    RETURNING id`
 
	var id int64 
    err := db.DB.QueryRow(query, e.ID, userId).Scan(&id)
    
	return err

}

func (e Event) Unregister(userId int64, eventId int64) error{
	query := `
    DELETE FROM registrations WHERE event_id = $1 AND user_id = $2`
 
    preparedStatement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer preparedStatement.Close()

	_ , err = preparedStatement.Exec(eventId, userId)
	return err

}
	