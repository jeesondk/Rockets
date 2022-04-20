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
	resp := h.HandleLaunchMessage(message)

	assert.Equal(t, expected, resp)

}
