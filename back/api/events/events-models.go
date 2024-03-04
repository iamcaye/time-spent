package events

type Event struct {
	Id       int
	Name     string
	Category string
	Date     string
}

type Events []Event
