package controllers

import "sync"

type Controller struct {
	Mutex sync.Mutex
}

func NewController() *Controller {
	return &Controller{}
}
