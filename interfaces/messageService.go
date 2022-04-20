package interfaces

import (
	"RocketService/dto"
	"RocketService/entities"
)

type MessageService interface {
	HandleMessage(metadata DTO.MetaData, message interface{}) (entities.Rocket, error)
	HandleLaunchMessage(data interface{}) (DTO.RocketLaunched, error)
	HandleSpeedIncreaseMessage(data interface{}) (DTO.RocketSpeedIncreased, error)
	HandleSpeedDecreaseMessage(data interface{}) (DTO.RocketSpeedDecreased, error)
	HandleMissionChangedMessage(data interface{}) (DTO.RocketMissionChanged, error)
	HandleRocketExplodedMessage(data interface{}) (DTO.RocketExploded, error)
}
