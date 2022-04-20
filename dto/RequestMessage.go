package DTO

import "RocketService/enum"

type RequestMessage struct {
	Metadata struct {
		Channel       string           `json:"channel"`
		MessageNumber int64            `json:"messageNumber"`
		MessageTime   string           `json:"messageTime"`
		MessageType   enum.MessageType `json:"messageType"`
	} `json:"metadata"`
	Message interface{} `json:"message"`
}
