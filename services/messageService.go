package services

import (
	DTO "RocketService/dto"
	"fmt"
)

type MessageService struct {
}

func (m *MessageService) HandleLaunchMessage(data interface{}) (DTO.RocketLaunched, error) {
	d := data.(map[string]interface{})

	msg := DTO.RocketLaunched{}
	errMsg := ""

	if rocketType, ok := d["type"].(string); ok {
		msg.Type = rocketType
	} else {
		errMsg += "Unable to parse type. "
	}

	if mission, ok := d["mission"].(string); ok {
		msg.Mission = mission
	} else {
		errMsg += "Unable to parse mission. "
	}

	if speed, ok := d["launchSpeed"].(int); ok {
		if speed < 0 {
			speed = speed * -1
		}
		msg.LaunchSpeed = speed

	} else {
		errMsg += "Unable to parse launchSpeed. "
	}

	if errMsg != "" {
		err := fmt.Errorf("caught errors while parsing data: %s", errMsg)
		return msg, err
	}

	return msg, nil
}

func (m *MessageService) HandleSpeedIncreaseMessage(data interface{}) (DTO.RocketSpeedIncreased, error) {
	d := data.(map[string]interface{})

	msg := DTO.RocketSpeedIncreased{}
	errMsg := ""

	if speed, ok := d["by"].(int); ok {
		if speed < 0 {
			speed = speed * -1
		}
		msg.By = speed
	} else {
		errMsg += "Unable to parse speed. "
	}

	if errMsg != "" {
		err := fmt.Errorf("caught errors while parsing data: %s", errMsg)
		return msg, err
	}

	return msg, nil
}

func (m *MessageService) HandleSpeedDecreaseMessage(data interface{}) (DTO.RocketSpeedDecreased, error) {
	d := data.(map[string]interface{})

	msg := DTO.RocketSpeedDecreased{}
	errMsg := ""

	if speed, ok := d["by"].(int); ok {
		if speed < 0 {
			speed = speed * -1
		}
		msg.By = speed
	} else {
		errMsg += "Unable to parse speed. "
	}

	if errMsg != "" {
		err := fmt.Errorf("caught errors while parsing data: %s", errMsg)
		return msg, err
	}

	return msg, nil
}

func (m *MessageService) HandleMissionChangedMessage(data interface{}) (DTO.RocketMissionChanged, error) {
	return DTO.RocketMissionChanged{}, nil
}

func (m *MessageService) HandleRocketExploedeMessage(data interface{}) (DTO.RocketExploded, error) {
	return DTO.RocketExploded{}, nil
}

func NewMessageHandler() *MessageService {
	return &MessageService{}
}
