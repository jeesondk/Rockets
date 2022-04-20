package services

import (
	DTO "RocketService/dto"
)

type MessageHandler struct {
}

func (m *MessageHandler) HandleLaunchMessage(data interface{}) DTO.RocketLaunched {
	d := data.(map[string]interface{})

	msg := DTO.RocketLaunched{}

	if rocketType, ok := d["type"].(string); ok {
		msg.Type = rocketType
	}
	if mission, ok := d["mission"].(string); ok {
		msg.Mission = mission
	}

	if launchSpeed, ok := d["launchSpeed"].(int); ok {
		msg.LaunchSpeed = launchSpeed
	}

	return msg
}

func NewMessageHandler() *MessageHandler {
	return &MessageHandler{}
}
