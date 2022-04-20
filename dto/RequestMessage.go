package DTO

import "RocketService/enum"

type RequestMessage struct {
	Metadata MetaData    `json:"metadata"`
	Message  interface{} `json:"message"`
}

type MetaData struct {
	Channel       string           `json:"channel"`
	MessageNumber int64            `json:"messageNumber"`
	MessageTime   string           `json:"messageTime"`
	MessageType   enum.MessageType `json:"messageType"`
}
