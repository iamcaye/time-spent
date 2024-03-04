package events

import (
	"fmt"
)

func (e *Event) String() string {
	return fmt.Sprintf("Event: %d %s %s %s", e.Id, e.Name, e.Category, e.Date)
}

func NewEvent(id int, name, category, date string) *Event {
	return &Event{id, name, category, date}
}

func SaveEvent(e *Event) (err error) {
	fmt.Println("Event saved:", e)
	return nil
}

func GetEvent(id int) (e *Event, err error) {
	return NewEvent(id, "Test Event", "Test Category", "Test Date"), nil
}

func GetAllEvents() (events Events, err error) {
	return Events{Event{1, "Test Event", "Test Category", "Test Date"}}, nil
}
