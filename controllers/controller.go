package controllers

import (
	"RocketService/interfaces"
	"RocketService/services"
)

type Controller struct {
	MessageService interfaces.MessageService
}

func NewController() *Controller {
	return &Controller{
		MessageService: services.NewMessageService(),
	}
}
