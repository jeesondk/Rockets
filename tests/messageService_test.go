package tests

import (
	DTO "RocketService/dto"
	"RocketService/entities"
	"RocketService/services"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
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

	h := services.NewMessageService()
	resp, err := h.HandleLaunchMessage(message)

	assert.Nil(t, err)
	assert.Equal(t, expected, resp)

}

func TestCanHandleLaunchNegSpeedMsg(t *testing.T) {
	message := make(map[string]interface{})
	message["type"] = "Falcon-9"
	message["launchSpeed"] = -500
	message["mission"] = "ARTEMIS"

	expected := DTO.RocketLaunched{
		Type:        "Falcon-9",
		LaunchSpeed: 500,
		Mission:     "ARTEMIS",
	}

	h := services.NewMessageService()
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

	h := services.NewMessageService()
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

	h := services.NewMessageService()
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

	h := services.NewMessageService()
	resp, err := h.HandleLaunchMessage(message)

	assert.Error(t, err, "caught errors while parsing data: Unable to parse launchSpeed. ")
	assert.Equal(t, expected, resp)

}

func TestCanHandleSpeedIncreaseMsg(t *testing.T) {
	message := make(map[string]interface{})
	message["by"] = 2500

	expected := DTO.RocketSpeedIncreased{
		By: 2500,
	}

	h := services.NewMessageService()
	resp, err := h.HandleSpeedIncreaseMessage(message)

	assert.Nil(t, err)
	assert.Equal(t, expected, resp)

}

func TestCanHandleSpeedDecreaseMsg(t *testing.T) {
	message := make(map[string]interface{})
	message["by"] = 1500

	expected := DTO.RocketSpeedDecreased{
		By: 1500,
	}

	h := services.NewMessageService()
	resp, err := h.HandleSpeedDecreaseMessage(message)

	assert.Nil(t, err)
	assert.Equal(t, expected, resp)

}

func TestCanHandleNegSpeedIncreaseMsg(t *testing.T) {
	message := make(map[string]interface{})
	message["by"] = -2500

	expected := DTO.RocketSpeedIncreased{
		By: 2500,
	}

	h := services.NewMessageService()
	resp, err := h.HandleSpeedIncreaseMessage(message)

	assert.Nil(t, err)
	assert.Equal(t, expected, resp)

}

func TestCanHandleNegSpeedDecreaseMsg(t *testing.T) {
	message := make(map[string]interface{})
	message["by"] = -1500

	expected := DTO.RocketSpeedDecreased{
		By: 1500,
	}

	h := services.NewMessageService()
	resp, err := h.HandleSpeedDecreaseMessage(message)

	assert.Nil(t, err)
	assert.Equal(t, expected, resp)

}

func TestFailHandleSpeedIncreaseMsg(t *testing.T) {
	message := make(map[string]interface{})
	message["by"] = "1500"

	expected := DTO.RocketSpeedIncreased{}

	h := services.NewMessageService()
	resp, err := h.HandleSpeedIncreaseMessage(message)

	assert.Error(t, err, "caught errors while parsing data: Unable to parse speed. ")
	assert.Equal(t, expected, resp)

}

func TestFailHandleSpeedDecreaseMsg(t *testing.T) {
	message := make(map[string]interface{})
	message["by"] = "1500"

	expected := DTO.RocketSpeedDecreased{}

	h := services.NewMessageService()
	resp, err := h.HandleSpeedDecreaseMessage(message)

	assert.Error(t, err, "caught errors while parsing data: Unable to parse speed. ")
	assert.Equal(t, expected, resp)

}

func TestCanHandleNewMissionMsg(t *testing.T) {
	message := make(map[string]interface{})
	message["newMission"] = "SHUTTLE_MIR"

	expected := DTO.RocketMissionChanged{
		NewMission: "SHUTTLE_MIR",
	}

	h := services.NewMessageService()
	resp, err := h.HandleMissionChangedMessage(message)

	assert.Nil(t, err)
	assert.Equal(t, expected, resp)

}

func TestFailHandleNewMissionMsg(t *testing.T) {
	message := make(map[string]interface{})
	message["newMission"] = 42

	expected := DTO.RocketMissionChanged{}

	h := services.NewMessageService()
	resp, err := h.HandleMissionChangedMessage(message)

	assert.Error(t, err, "caught errors while parsing data: Unable to parse mission. ")
	assert.Equal(t, expected, resp)

}

func TestCanHandleExplodedMsg(t *testing.T) {
	message := make(map[string]interface{})
	message["reason"] = "PRESSURE_VESSEL_FAILURE"

	expected := DTO.RocketExploded{
		Reason: "PRESSURE_VESSEL_FAILURE",
	}

	h := services.NewMessageService()
	resp, err := h.HandleRocketExplodedMessage(message)

	assert.Nil(t, err)
	assert.Equal(t, expected, resp)

}

func TestFailHandleExplodedMsg(t *testing.T) {
	message := make(map[string]interface{})
	message["reason"] = 42

	expected := DTO.RocketExploded{}

	h := services.NewMessageService()
	resp, err := h.HandleRocketExplodedMessage(message)

	assert.Error(t, err, "caught errors while parsing data: Unable to parse reason. ")
	assert.Equal(t, expected, resp)

}

func TestMessageHandleMessages(t *testing.T) {
	rocket := entities.Rocket{
		ID:   "1",
		Type: "Falcon 9",
		Speed: entities.RocketSpeed{
			Current: 0,
			Max:     10000,
		},
		Mission:    "ARTEMIS",
		LaunchDate: time.Now(),
	}

	tests := []struct {
		name     string
		rocket   entities.Rocket
		metadata DTO.MetaData
		message  map[string]interface{}
		expected entities.Rocket
	}{
		{"CAN_HANDLE_SPEED_INCREASE_MSG",
			rocket,
			DTO.MetaData{
				Channel:       "rocket_id",
				MessageType:   "SPEED_INCREASE",
				MessageNumber: 1,
				MessageTime:   time.Now().String(),
			},
			map[string]interface{}{
				"by": 2500,
			},
			entities.Rocket{
				ID: "rocket_id",
				Speed: entities.RocketSpeed{
					Current: 3500,
					Max:     10000,
				},
			},
		},
		{"CAN_HANDLE_SPEED_DECREASE_MSG",
			rocket,
			DTO.MetaData{
				Channel:       "rocket_id",
				MessageType:   "SPEED_DECREASE",
				MessageNumber: 1,
				MessageTime:   time.Now().String(),
			},
			map[string]interface{}{
				"by": -1500,
			},
			entities.Rocket{
				ID: "rocket_id",
				Speed: entities.RocketSpeed{
					Current: 3500,
					Max:     10000,
				},
			},
		},
		{"CAN_HANDLE_MISSION_CHANGED_MSG",
			rocket,
			DTO.MetaData{
				Channel:       "rocket_id",
				MessageType:   "MISSION_CHANGED",
				MessageNumber: 1,
				MessageTime:   time.Now().String(),
			},
			map[string]interface{}{
				"newMission": "SHUTTLE_MIR",
			},
			entities.Rocket{
				ID:      "rocket_id",
				Mission: "SHUTTLE_MIR",
			},
		},
		{"CAN_HANDLE_EXPLODED_MSG",
			rocket,
			DTO.MetaData{
				Channel:       "rocket_id",
				MessageType:   "EXPLODED",
				MessageNumber: 1,
				MessageTime:   time.Now().String(),
			},
			map[string]interface{}{
				"reason": "PRESSURE_VESSEL_FAILURE",
			},
			entities.Rocket{
				ID:     "rocket_id",
				Status: entities.RocketStatus{Active: false, Reason: "EXPLODED"},
			},
		},
	}

	for i, tc := range tests {
		t.Run(fmt.Sprintf("Test %d: %s", i, tc.name), func(t *testing.T) {
			s := services.NewMessageService()
			resp, err := s.HandleMessage(tc.metadata, tc.message)

			assert.Nil(t, err)
			assert.Equal(t, tc.expected, resp)
		})
	}
}
