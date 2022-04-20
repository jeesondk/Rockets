package tests

import (
	DTO "RocketService/dto"
	"RocketService/entities"
	"github.com/stretchr/testify/mock"
)

type messageServiceMock struct {
	mock.Mock
}

func (m *messageServiceMock) HandleMessage(metadata DTO.MetaData, message interface{}) (entities.Rocket, error) {
	args := m.Called(metadata, message)
	return args.Get(0).(entities.Rocket), args.Error(1)
}

func (m *messageServiceMock) HandleLaunchMessage(data interface{}) (DTO.RocketLaunched, error) {
	args := m.Called(data)
	return args.Get(0).(DTO.RocketLaunched), args.Error(1)
}

func (m *messageServiceMock) HandleSpeedIncreaseMessage(data interface{}) (DTO.RocketSpeedIncreased, error) {
	args := m.Called(data)
	return args.Get(0).(DTO.RocketSpeedIncreased), args.Error(1)
}

func (m *messageServiceMock) HandleSpeedDecreaseMessage(data interface{}) (DTO.RocketSpeedDecreased, error) {
	args := m.Called(data)
	return args.Get(0).(DTO.RocketSpeedDecreased), args.Error(1)
}

func (m *messageServiceMock) HandleMissionChangedMessage(data interface{}) (DTO.RocketMissionChanged, error) {
	args := m.Called(data)
	return args.Get(0).(DTO.RocketMissionChanged), args.Error(1)
}

func (m *messageServiceMock) HandleRocketExplodedMessage(data interface{}) (DTO.RocketExploded, error) {
	args := m.Called(data)
	return args.Get(0).(DTO.RocketExploded), args.Error(1)
}
