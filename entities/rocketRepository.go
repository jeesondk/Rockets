package entities

type RocketRepository interface {
	GetAll() ([]Rocket, error)
	GetByID(id string) (Rocket, error)
	Create(rocket Rocket) (Rocket, error)
	Update(rocket Rocket) (Rocket, error)
}

type Rockets struct {
	rocketCollection map[string]Rocket
}

func (r Rockets) GetAll() ([]Rocket, error) {
	//TODO implement me
	panic("implement me")
}

func (r Rockets) GetByID(id string) (Rocket, error) {
	//TODO implement me
	panic("implement me")
}

func (r Rockets) Create(rocket Rocket) (Rocket, error) {
	//TODO implement me
	panic("implement me")
}

func (r Rockets) Update(rocket Rocket) (Rocket, error) {
	//TODO implement me
	panic("implement me")
}

func NewRocketRepository() RocketRepository {
	return &Rockets{
		rocketCollection: make(map[string]Rocket),
	}
}
