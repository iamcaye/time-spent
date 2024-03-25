package events

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type EventService struct {
	db *sqlx.DB
}

func (srv *EventService) GetEvents() []Event {
	events := []Event{}
	err := srv.db.Select(&events, "SELECT * FROM events")
	if err != nil {
		panic(err)
	}

	return events
}

func (srv *EventService) GetEventByID(id int) Event {
	// This function will return an event by ID
	// For now, it will return a hardcoded event
	event := Event{}
	err := srv.db.Get(&event, "SELECT * FROM events WHERE id = ?", id)
	if err != nil {
		panic(err)
	}

	return event
}

func (eventService *EventService) CreateEvent(event Event) Event {
	// This function will create a new event
	err := eventService.db.Get(&event,
		"INSERT INTO events (name, id_user, id_category) VALUES (?, ?, ?) RETURNING *",
		event.Name, event.IdUser, event.IdCategory)

	if err != nil {
		panic(err)
	}

	return event
}

func (eventService *EventService) UpdateEvent(event Event) bool {
	// This function will update an existing event
	fmt.Println(event)
	res, err := eventService.db.Exec(
		"UPDATE events SET name = ?, id_user = ?, id_category = ? WHERE id = ?",
		event.Name, event.IdUser, event.IdCategory, event.ID)

	if err != nil {
		panic(err)
	}

	n, _ := res.RowsAffected()
	fmt.Println(n)
	return n > 0
}

func (eventService *EventService) DeleteEvent(id int) bool {
	// This function will delete an existing event
	res, err := eventService.db.Exec("DELETE FROM events WHERE id = ?", id)
	if err != nil {
		panic(err)
	}

	n, _ := res.RowsAffected()
	return n > 0
}

func (srv *EventService) StartEvent(id int) bool {
	// This function will start an event
	res, err := srv.db.Exec("INSERT INTO events_timestamps (id_event) VALUES (?)", id)
	if err != nil {
		panic(err)
	}

	n, _ := res.RowsAffected()
	return n > 0
}

func (srv *EventService) StopEvent(id int) bool {
	res, err := srv.db.Exec("UPDATE events_timestamps SET end_time = NOW() WHERE id_event = ? AND end_time IS NULL", id)
	if err != nil {
		panic(err)
	}

	n, _ := res.RowsAffected()
	return n > 0
}
