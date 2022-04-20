package entities

type Rocket struct {
	ID         string
	Type       string
	Speed      RocketSpeed
	Mission    string
	Status     RocketStatus
	LaunchDate string
}
