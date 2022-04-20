package tests

import (
	"RocketService/entities"
	"github.com/stretchr/testify/mock"
)

type rocketRepositoryMock struct {
	mock.Mock
}

func (r *rocketRepositoryMock) GetAll() ([]entities.Rocket, error) {
	args := r.Called()
	return args.Get(0).([]entities.Rocket), args.Error(1)
}
func (r *rocketRepositoryMock) GetByID(id string) (entities.Rocket, error) {
	args := r.Called(id)
	return args.Get(0).(entities.Rocket), args.Error(1)
}
func (r *rocketRepositoryMock) Create(rocket entities.Rocket) (entities.Rocket, error) {
	args := r.Called(rocket)
	return args.Get(0).(entities.Rocket), args.Error(1)
}
func (r *rocketRepositoryMock) Update(rocket entities.Rocket) (entities.Rocket, error) {
	args := r.Called(rocket)
	return args.Get(0).(entities.Rocket), args.Error(1)
}
