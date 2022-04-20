package enum

type MessageType string

const (
	RocketLaunched       = "RocketLaunched"
	RocketSpeedIncreased = "RocketSpeedIncreased"
	RocketSpeedDecreased = "RocketSpeedDecreased"
	RocketExploded       = "RocketExploded"
	RocketMissionChanged = "RocketMissionChanged"
)
