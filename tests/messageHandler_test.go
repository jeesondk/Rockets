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

func TestCanHandleSpeedIncreaseMsg(t *testing.T) {
	message := make(map[string]interface{})
	message["by"] = 2500

	expected := DTO.RocketSpeedIncreased{
		By: 2500,
	}

	h := services.NewMessageHandler()
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

	h := services.NewMessageHandler()
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

	h := services.NewMessageHandler()
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

	h := services.NewMessageHandler()
	resp, err := h.HandleSpeedDecreaseMessage(message)

	assert.Nil(t, err)
	assert.Equal(t, expected, resp)

}

func TestFailHandleSpeedIncreaseMsg(t *testing.T) {
	message := make(map[string]interface{})
	message["by"] = "1500"

	expected := DTO.RocketSpeedIncreased{}

	h := services.NewMessageHandler()
	resp, err := h.HandleSpeedIncreaseMessage(message)

	assert.Error(t, err, "caught errors while parsing data: Unable to parse speed. ")
	assert.Equal(t, expected, resp)

}

func TestFailHandleSpeedDecreaseMsg(t *testing.T) {
	message := make(map[string]interface{})
	message["by"] = "1500"

	expected := DTO.RocketSpeedDecreased{}

	h := services.NewMessageHandler()
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

	h := services.NewMessageHandler()
	resp, err := h.HandleMissionChangedMessage(message)

	assert.Nil(t, err)
	assert.Equal(t, expected, resp)

}

func TestFailHandleNewMissionMsg(t *testing.T) {
	message := make(map[string]interface{})
	message["newMission"] = 42

	expected := DTO.RocketMissionChanged{}

	h := services.NewMessageHandler()
	resp, err := h.HandleMissionChangedMessage(message)

	assert.Error(t, err, "caught errors while parsing data: Unable to parse mission. ")
	assert.Equal(t, expected, resp)

}

func TestCanHandleExplodedMsg(t *testing.T) {
	message := make(map[string]interface{})
	message["reason"] = "SHUTTLE_MIR"

	expected := DTO.RocketExploded{
		Reason: "PRESSURE_VESSEL_FAILURE",
	}

	h := services.NewMessageHandler()
	resp, err := h.HandleRocketExplodedMessage(message)

	assert.Nil(t, err)
	assert.Equal(t, expected, resp)

}

func TestFailHandleExplodedMsg(t *testing.T) {
	message := make(map[string]interface{})
	message["reason"] = 42

	expected := DTO.RocketExploded{}

	h := services.NewMessageHandler()
	resp, err := h.HandleRocketExplodedMessage(message)

	assert.Error(t, err, "caught errors while parsing data: Unable to parse reason. ")
	assert.Equal(t, expected, resp)

}
