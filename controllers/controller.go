package controllers

import (
	"RocketService/services"
)

type Controller struct {
	MessageService services.MESSAGESERVICE
}

func NewController() *Controller {
	return &Controller{
		MessageService: services.NewMessageService(),
	}
}
