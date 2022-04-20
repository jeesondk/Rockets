package services

import (
	"RocketService/entities"
)

type RocketService struct {
	RocketRepository entities.RocketRepository
}

type ROCKETSERVICE interface {
	GetRocket(id string) (entities.Rocket, error)
	GetAllRockets() ([]entities.Rocket, error)
	CreateRocket(rocket entities.Rocket) (string, error)
	UpdateRocket(id string, rocket entities.Rocket) (string, error)
}

func (r *RocketService) GetRocket(id string) (entities.Rocket, error) {

	return entities.Rocket{}, nil
}

func (r *RocketService) GetAllRockets() ([]entities.Rocket, error) {
	return []entities.Rocket{}, nil
}

func (r *RocketService) CreateRocket(rocket entities.Rocket) (string, error) {
	return rocket.ID, nil
}

func (r *RocketService) UpdateRocket(id string, rocket entities.Rocket) (string, error) {
	return id, nil
}

func NewRocketService() ROCKETSERVICE {
	return new(RocketService)
}
