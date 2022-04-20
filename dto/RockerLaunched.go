package DTO

type RocketLaunched struct {
	LaunchSpeed int64  `json:"launchSpeed"`
	Mission     string `json:"mission"`
	Type        string `json:"type"`
}
