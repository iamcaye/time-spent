package events

type Event struct {
	ID          int
	Name        string
	Description string
	Date        string
}

type EventService struct{}

func (eventService *EventService) GetEvents() []Event {
	// This function will return a list of events
	// For now, it will return a list of hardcoded events
	events := []Event{
		{ID: 1, Name: "Event 1", Description: "The first event", Date: "01/01/2021"},
		{ID: 2, Name: "Event 2", Description: "The second event", Date: "02/01/2021"},
		{ID: 3, Name: "Event 3", Description: "The third event", Date: "03/01/2021"},
	}
	return events
}

func (eventService *EventService) GetEventByID(id int) Event {
	// This function will return an event by ID
	// For now, it will return a hardcoded event
	event := Event{ID: 1, Name: "Event 1", Description: "The first event", Date: "01/01/2021"}
	return event
}

func (eventService *EventService) CreateEvent(event Event) Event {
	// This function will create a new event
	// For now, it will return the event that was created
	return event
}

func (eventService *EventService) UpdateEvent(event Event) Event {
	// This function will update an existing event
	// For now, it will return the event that was updated
	return event
}

func (eventService *EventService) DeleteEvent(id int) bool {
	// This function will delete an event by ID
	// For now, it will do nothing
	return true
}
