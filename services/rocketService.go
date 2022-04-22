package services

import (
	DTO "RocketService/dto"
	"RocketService/entities"
	"RocketService/enum"
	"sync"
)

type RockerServiceInterface interface {
	GetRocket(id string) (*entities.Rocket, error)
	GetAllRockets() (*[]entities.Rocket, error)
	CreateRocket(rocket *entities.Rocket) error
	UpdateRockets()
}

type RocketService struct {
	RocketRepository *entities.RocketRepository
	Mux              *sync.Mutex
	Wg               *sync.WaitGroup
	MessageService   *MessageService
}

func (r *RocketService) GetRocket(id string) (*entities.Rocket, error) {

	r.Wg.Add(1)
	rocket, err := r.RocketRepository.GetByID(id, r.Mux, r.Wg)
	r.Wg.Wait()
	if err != nil {
		return new(entities.Rocket), err
	}
	return rocket, nil
}

func (r *RocketService) GetAllRockets() (*[]entities.Rocket, error) {
	r.Wg.Add(1)
	rockets := r.RocketRepository.GetAllRockets(r.Mux, r.Wg)
	r.Wg.Wait()
	return rockets, nil
}

func (r *RocketService) CreateRocket(rocket *entities.Rocket) error {
	r.Wg.Add(1)
	err := r.RocketRepository.Create(rocket, r.Mux, r.Wg)
	r.Wg.Wait()
	if err != nil {
		return err
	}
	return nil
}

func (r *RocketService) UpdateRockets() {
	r.Wg.Add(1)
	rockets := r.RocketRepository.GetAllRockets(r.Mux, r.Wg)
	r.Wg.Wait()

	for _, rocket := range *rockets {

		r.Wg.Add(1)
		events, _ := r.RocketRepository.GetEvent(rocket.ID, r.Mux, r.Wg)
		r.Wg.Wait()

		shouldUpdate := true
		offset := rocket.EventCursor + 1
		j := 0

		for i := offset; i <= len(events); i++ {
			if i != events[j].EventId {
				shouldUpdate = false
			}
			j++
		}
		if shouldUpdate && events != nil {
			r.rollUpEvents(events, rocket.ID)
		}

	}
}

func (r *RocketService) rollUpEvents(events []entities.RocketEvent, id string) error {
	r.Wg.Add(1)
	rocket, err := r.RocketRepository.GetByID(id, r.Mux, r.Wg)
	r.Wg.Wait()
	if err != nil {
		return err
	}

	errs := make([]error, 0)

	for _, e := range events {

		if err != nil {
			errs = append(errs, err)
		}

		switch e.EventType {
		case enum.RocketSpeedIncreased:
			specificEvent, errHandle := r.MessageService.HandleSpeedIncreaseMessage(e.Event)
			if errHandle != nil {
				return errHandle
			}

			err := r.speedIncreaseMessage(rocket, specificEvent)
			if err != nil {
				return err
			}

		case enum.RocketSpeedDecreased:
			specificEvent, errHandle := r.MessageService.HandleSpeedDecreaseMessage(e.Event)
			if errHandle != nil {
				return errHandle
			}

			err := r.speedDecreaseMessage(rocket, specificEvent)
			if err != nil {
				return err
			}

		case enum.RocketMissionChanged:
			specificEvent, errHandle := r.MessageService.HandleMissionChangedMessage(e.Event)
			if errHandle != nil {
				return errHandle
			}

			err := r.missionChangedMessage(rocket, specificEvent)
			if err != nil {
				return err
			}

		case enum.RocketExploded:
			specificEvent, errHandle := r.MessageService.HandleRocketExplodedMessage(e.Event)
			if errHandle != nil {
				return errHandle
			}

			err := r.rocketExplodedMessage(rocket, specificEvent)
			if err != nil {
				return err
			}
		}
		rocket.EventCursor++
	}
	return nil
}

func (r *RocketService) speedIncreaseMessage(rocket *entities.Rocket, event DTO.RocketSpeedIncreased) error {

	rocket.Speed.Update(event.By)

	r.Wg.Add(1)
	err := r.RocketRepository.Update(rocket, r.Mux, r.Wg)
	r.Wg.Wait()

	if err != nil {
		return err
	}
	return nil
}

func (r *RocketService) speedDecreaseMessage(rocket *entities.Rocket, event DTO.RocketSpeedDecreased) error {
	rocket.Speed.Update(event.By * -1)

	r.Wg.Add(1)
	err := r.RocketRepository.Update(rocket, r.Mux, r.Wg)
	r.Wg.Wait()

	if err != nil {
		return err
	}
	return nil
}

func (r *RocketService) missionChangedMessage(rocket *entities.Rocket, event DTO.RocketMissionChanged) error {
	rocket.Mission = event.NewMission

	r.Wg.Add(1)
	err := r.RocketRepository.Update(rocket, r.Mux, r.Wg)
	r.Wg.Wait()

	if err != nil {
		return err
	}
	return nil
}

func (r *RocketService) rocketExplodedMessage(rocket *entities.Rocket, event DTO.RocketExploded) error {
	rocket.Status.Active = false
	rocket.Status.Reason = event.Reason

	r.Wg.Add(1)
	err := r.RocketRepository.Update(rocket, r.Mux, r.Wg)
	r.Wg.Wait()

	if err != nil {
		return err
	}
	return nil
}

func NewRocketService() RocketService {
	return RocketService{}
}
