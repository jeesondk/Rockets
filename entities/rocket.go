package entities

type Rocket struct {
	ID          string
	RocketType  string
	Speed       RocketSpeed
	Mission     string
	Status      RocketStatus
	LaunchDate  string
	EventCursor int
}
