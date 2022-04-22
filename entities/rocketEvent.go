package entities

type RocketEvent struct {
	RocketId  string
	EventId   int
	EventType string
	Event     interface{}
}
