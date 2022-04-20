package tests

import (
	DTO "RocketService/dto"
	"RocketService/services"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCanHandleLaunchMsg(t *testing.T) {
	requestMsg := DTO.RequestMessage{
		Metadata: DTO.MetaData{
			Channel:       "193270a9-c9cf-404a-8f83-838e71d9ae67",
			MessageNumber: 1,
			MessageTime:   "2022-02-02T19:39:05.86337+01:00",
			MessageType:   "RocketLaunched",
		},
		Message: `{
					"type": "Falcon-9",
					"launchSpeed": 500,
					"mission": "ARTEMIS"  
				}`,
	}

	expected := DTO.RocketLaunched{
		Type:        "Falcon-9",
		LaunchSpeed: 500,
		Mission:     "ARTEMIS",
	}

	h := services.NewMessageHandler()
	resp, err := h.HandleLaunchMessage(requestMsg)

	assert.Equal(t, nil, err)
	assert.Equal(t, expected, resp)

}
