package services

import (
	DTO "RocketService/dto"
	"fmt"
)

type MessageHandler struct {
}

func (m *MessageHandler) HandleLaunchMessage(data interface{}) (DTO.RocketLaunched, error) {
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

	if launchSpeed, ok := d["launchSpeed"].(int); ok {
		msg.LaunchSpeed = launchSpeed
	} else {
		errMsg += "Unable to parse launchSpeed. "
	}

	if errMsg != "" {
		err := fmt.Errorf("caught errors while parsing data: %s", errMsg)
		return msg, err
	}

	return msg, nil
}

func (m *MessageHandler) HandleSpeedIncreaseMessage(data interface{}) (DTO.RocketSpeedIncreased, error) {
	d := data.(map[string]interface{})

	msg := DTO.RocketSpeedIncreased{}
	errMsg := ""

	if speed, ok := d["by"].(int); ok {
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

func (m *MessageHandler) HandleSpeedDecreaseMessage(data interface{}) (DTO.RocketSpeedDecreased, error) {
	d := data.(map[string]interface{})

	msg := DTO.RocketSpeedDecreased{}
	errMsg := ""

	if speed, ok := d["by"].(int); ok {
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

func NewMessageHandler() *MessageHandler {
	return &MessageHandler{}
}
