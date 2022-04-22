package entities

import (
	DTO "RocketService/dto"
	"fmt"
	"sync"
)

type RockerRepository interface {
	GetAll(ch chan *[]Rocket, mtx *sync.Mutex, wg *sync.WaitGroup)
	GetByID(id string, ch chan *Rocket, mtx *sync.Mutex, wg *sync.WaitGroup, errCh chan error)
	Create(rocket Rocket, ch chan *Rocket, mtx *sync.Mutex, wg *sync.WaitGroup, errCh chan error)
	Update(rocket Rocket, ch chan *Rocket, mtx *sync.Mutex, wg *sync.WaitGroup, errCh chan error)
}

type RocketRepository struct {
	rocketCollection map[string]Rocket
	//rocketEventCollection map[string]RocketEvent
	rocketEventCollection map[string][]RocketEvent
}

func (r RocketRepository) GetAllRockets(mtx *sync.Mutex, wg *sync.WaitGroup) *[]Rocket {
	mtx.Lock()
	rockets := make([]Rocket, 0)
	for _, value := range r.rocketCollection {
		rockets = append(rockets, value)
	}
	mtx.Unlock()
	wg.Done()
	return &rockets
}

func (r RocketRepository) GetByID(id string, mtx *sync.Mutex, wg *sync.WaitGroup) (*Rocket, error) {

	mtx.Lock()
	rocket, ok := r.rocketCollection[id]
	mtx.Unlock()

	if !ok {
		return &rocket, fmt.Errorf("rocket not found")
	}
	wg.Done()
	return &rocket, nil
}

func (r RocketRepository) Create(rocket *Rocket, mtx *sync.Mutex, wg *sync.WaitGroup) error {
	mtx.Lock()
	id := rocket.ID
	_, exists := r.rocketCollection[id]
	mtx.Unlock()

	if exists {
		return fmt.Errorf("rocket already exists")
	}

	mtx.Lock()
	r.rocketCollection[id] = *rocket
	mtx.Unlock()

	wg.Done()
	return nil
}

func (r RocketRepository) Update(rocket *Rocket, mtx *sync.Mutex, wg *sync.WaitGroup) error {
	mtx.Lock()
	defer mtx.Unlock()

	var err error
	res, ok := r.rocketCollection[rocket.ID]
	if !ok {
		err = fmt.Errorf("rocket not found")
	}

	if !res.Status.Active {
		err = fmt.Errorf("rocket is not active, reason: %s", res.Status.Reason)
	}

	r.rocketCollection[rocket.ID] = *rocket

	wg.Done()
	return err
}

func (r RocketRepository) GetEvent(rocketId string, mtx *sync.Mutex, wg *sync.WaitGroup) ([]RocketEvent, error) {
	mtx.Lock()
	res, ok := r.rocketEventCollection[rocketId]
	mtx.Unlock()

	if !ok {
		wg.Done()
		return nil, fmt.Errorf("rocket event not found")
	}
	wg.Done()

	return res, nil
}

func (r RocketRepository) AddEvent(metadata *DTO.MetaData, data interface{}, mtx *sync.Mutex, wg *sync.WaitGroup) error {
	rocketEvent := RocketEvent{
		metadata.Channel,
		metadata.MessageNumber,
		string(metadata.MessageType),
		data,
	}

	mtx.Lock()
	r.rocketEventCollection[metadata.Channel] = append(r.rocketEventCollection[metadata.Channel], rocketEvent)
	mtx.Unlock()

	wg.Done()
	return nil
}

func NewRocketRepository() RocketRepository {
	return RocketRepository{
		rocketCollection:      make(map[string]Rocket),
		rocketEventCollection: make(map[string][]RocketEvent),
	}
}
