package entities

import (
	"fmt"
	"sync"
)

type RocketRepository interface {
	GetAll(ch chan *[]Rocket)
	GetByID(id string, ch chan *Rocket, errCh chan error)
	Create(rocket Rocket, ch chan *Rocket, errCh chan error)
	Update(rocket Rocket, ch chan *Rocket, errCh chan error)
}

type Rockets struct {
	mux              sync.Mutex
	rocketCollection map[string]Rocket
}

func (r Rockets) GetAll(ch chan *[]Rocket) {
	defer close(ch)
	r.mux.Lock()
	rockets := make([]Rocket, len(r.rocketCollection))
	for _, value := range r.rocketCollection {
		rockets = append(rockets, value)
	}
	r.mux.Unlock()
	ch <- &rockets
}

func (r Rockets) GetByID(id string, ch chan *Rocket, errCh chan error) {
	defer close(ch)
	defer close(errCh)

	r.mux.Lock()
	var rocket = r.rocketCollection[id]
	r.mux.Unlock()

	if rocket == (Rocket{}) {
		errCh <- fmt.Errorf("Rocket not found")
		ch <- nil
	}

	errCh <- nil
	ch <- &rocket
}

func (r Rockets) Create(rocket Rocket, ch chan *Rocket, errCh chan error) {
	defer close(ch)
	defer close(errCh)

	r.mux.Lock()

	if r.rocketCollection[rocket.ID] != (Rocket{}) {
		r.mux.Unlock()
		errCh <- fmt.Errorf("Rocket already exists")
		ch <- nil
	}

	r.rocketCollection[rocket.ID] = rocket
	r.mux.Unlock()

	errCh <- nil
	ch <- &rocket
}

func (r Rockets) Update(rocket Rocket, ch chan *Rocket, errCh chan error) {
	defer close(ch)
	defer close(errCh)

	r.mux.Lock()

	res := r.rocketCollection[rocket.ID]
	if !res.Status.Active {
		r.mux.Unlock()
		errCh <- fmt.Errorf("Rocket is not active, reason: %s", res.Status.Reason)
		ch <- nil
	}

	if res == (Rocket{}) {
		r.mux.Unlock()
		errCh <- fmt.Errorf("Rocket not found")
		ch <- nil
	}
	r.rocketCollection[rocket.ID] = rocket
	r.mux.Unlock()

	errCh <- nil
	ch <- &res
}

func NewRocketRepository() *Rockets {
	return &Rockets{
		rocketCollection: make(map[string]Rocket),
	}
}
