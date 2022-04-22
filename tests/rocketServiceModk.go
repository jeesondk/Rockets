package tests

import (
	"RocketService/entities"
	"github.com/stretchr/testify/mock"
)

type rocketServiceMock struct {
	mock.Mock
}

func (m *rocketServiceMock) GetRocket(id string) (entities.Rocket, error) {
	args := m.Called(id)
	return args.Get(0).(entities.Rocket), args.Error(1)
}

func (m *rocketServiceMock) GetAllRockets() ([]entities.Rocket, error) {
	args := m.Called()
	return args.Get(0).([]entities.Rocket), args.Error(1)
}

func (m *rocketServiceMock) UpdateRockets() error {
	args := m.Called()
	return args.Error(0)
}
