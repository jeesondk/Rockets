package services

import (
	DTO "RocketService/dto"
	"fmt"
)

type MessageHandler struct {
}

func (m *MessageHandler) HandleLaunchMessage(message interface{}) (DTO.RocketLaunched, error) {
	return DTO.RocketLaunched{}, fmt.Errorf("Not implemented")
}

func NewMessageHandler() *MessageHandler {
	return &MessageHandler{}
}
