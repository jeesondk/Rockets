package tests

import (
	DTO "RocketService/dto"
	"RocketService/services"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCanHandleLaunchMsg(t *testing.T) {
	message := make(map[string]interface{})
	message["type"] = "Falcon-9"
	message["launchSpeed"] = 500
	message["mission"] = "ARTEMIS"

	expected := DTO.RocketLaunched{
		Type:        "Falcon-9",
		LaunchSpeed: 500,
		Mission:     "ARTEMIS",
	}

	h := services.NewMessageHandler()
	resp, err := h.HandleLaunchMessage(message)

	assert.Nil(t, err)
	assert.Equal(t, expected, resp)

}

func TestFailHandleLaunchMsgTypeErr(t *testing.T) {
	message := make(map[string]interface{})
	message["type"] = 99
	message["launchSpeed"] = 500
	message["mission"] = "ARTEMIS"

	expected := DTO.RocketLaunched{
		LaunchSpeed: 500,
		Mission:     "ARTEMIS",
	}

	h := services.NewMessageHandler()
	resp, err := h.HandleLaunchMessage(message)

	assert.Error(t, err, "caught errors while parsing data: Unable to parse type. ")
	assert.Equal(t, expected, resp)

}

func TestFailHandleLaunchMsgMissionErr(t *testing.T) {
	message := make(map[string]interface{})
	message["type"] = "Falcon-9"
	message["launchSpeed"] = 500
	message["mission"] = 42

	expected := DTO.RocketLaunched{
		Type:        "Falcon-9",
		LaunchSpeed: 500,
	}

	h := services.NewMessageHandler()
	resp, err := h.HandleLaunchMessage(message)

	assert.Error(t, err, "caught errors while parsing data: Unable to parse mission. ")
	assert.Equal(t, expected, resp)

}

func TestFailHandleLaunchMsgSpeedErr(t *testing.T) {
	message := make(map[string]interface{})
	message["type"] = "Falcon-9"
	message["launchSpeed"] = "500"
	message["mission"] = "ARTEMIS"

	expected := DTO.RocketLaunched{
		Type:    "Falcon-9",
		Mission: "ARTEMIS",
	}

	h := services.NewMessageHandler()
	resp, err := h.HandleLaunchMessage(message)

	assert.Error(t, err, "caught errors while parsing data: Unable to parse launchSpeed. ")
	assert.Equal(t, expected, resp)

}
