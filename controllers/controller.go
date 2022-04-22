package controllers

import (
	"RocketService/services"
)

type Controller struct {
	MessageService services.MessageServiceInterface
	RocketService  services.RockerServiceInterface
}

func NewController() Controller {
	return Controller{}
}
