package services

import (
	DTO "RocketService/dto"
	"RocketService/entities"
	"RocketService/enum"
	"RocketService/interfaces"
	"fmt"
)

type MessageService struct {
	Rockets entities.RocketRepository
}

func (m *MessageService) HandleLaunchMessage(data interface{}) (DTO.RocketLaunched, error) {
	d := data.(map[string]interface{})

	msg := DTO.RocketLaunched{}
	errMsg := ""

	if rocketType, ok := d["type"].(string); ok {
		msg.Type = rocketType
	} else {
		errMsg += "Unable to parse type. "
	}

	if mission, ok := d["mission"].(string); ok {
		msg.Mission = mission
	} else {
		errMsg += "Unable to parse mission. "
	}

	if speed, ok := d["launchSpeed"].(float64); ok {
		if speed < 0 {
			speed = speed * -1
		}
		msg.LaunchSpeed = speed

	} else {
		errMsg += "Unable to parse launchSpeed. "
	}

	if errMsg != "" {
		err := fmt.Errorf("caught errors while parsing data: %s", errMsg)
		return msg, err
	}

	return msg, nil
}

func (m *MessageService) HandleSpeedIncreaseMessage(data interface{}) (DTO.RocketSpeedIncreased, error) {
	d := data.(map[string]interface{})

	msg := DTO.RocketSpeedIncreased{}
	errMsg := ""

	if speed, ok := d["by"].(float64); ok {
		if speed < 0 {
			speed = speed * -1
		}
		msg.By = speed
	} else {
		errMsg += "Unable to parse speed. "
	}

	if errMsg != "" {
		err := fmt.Errorf("caught errors while parsing data: %s", errMsg)
		return msg, err
	}

	return msg, nil
}

func (m *MessageService) HandleSpeedDecreaseMessage(data interface{}) (DTO.RocketSpeedDecreased, error) {
	d := data.(map[string]interface{})

	msg := DTO.RocketSpeedDecreased{}
	errMsg := ""

	if speed, ok := d["by"].(float64); ok {
		if speed < 0 {
			speed = speed * -1
		}
		msg.By = speed
	} else {
		errMsg += "Unable to parse speed. "
	}

	if errMsg != "" {
		err := fmt.Errorf("caught errors while parsing data: %s", errMsg)
		return msg, err
	}

	return msg, nil
}

func (m *MessageService) HandleMissionChangedMessage(data interface{}) (DTO.RocketMissionChanged, error) {
	d := data.(map[string]interface{})

	msg := DTO.RocketMissionChanged{}
	errMsg := ""

	if mission, ok := d["newMission"].(string); ok {
		msg.NewMission = mission
	} else {
		errMsg += "Unable to parse mission. "
	}

	if errMsg != "" {
		err := fmt.Errorf("caught errors while parsing data: %s", errMsg)
		return msg, err
	}

	return msg, nil
}

func (m *MessageService) HandleRocketExplodedMessage(data interface{}) (DTO.RocketExploded, error) {
	d := data.(map[string]interface{})

	msg := DTO.RocketExploded{}
	errMsg := ""

	if exploded, ok := d["reason"].(string); ok {
		msg.Reason = exploded
	} else {
		errMsg += "Unable to parse reason. "
	}

	if errMsg != "" {
		err := fmt.Errorf("caught errors while parsing data: %s", errMsg)
		return msg, err
	}

	return msg, nil
}

func (m *MessageService) HandleMessage(metadata DTO.MetaData, data interface{}) (entities.Rocket, error) {
	var rocket entities.Rocket

	switch metadata.MessageType {
	case enum.RocketLaunched:
		msg, err := m.HandleLaunchMessage(data)
		if err != nil {
			return entities.Rocket{}, err
		}

		newRocket := entities.Rocket{
			ID:      metadata.Channel,
			Type:    msg.Type,
			Mission: msg.Mission,
			Speed: entities.RocketSpeed{
				Current: msg.LaunchSpeed,
				Max:     msg.LaunchSpeed,
			},
			LaunchDate: metadata.MessageTime,
			Status: entities.RocketStatus{
				Active: true,
				Reason: "",
			},
		}

		ch := make(chan *entities.Rocket)
		errCh := make(chan error)
		go m.Rockets.Create(newRocket, ch, errCh)

		rocket = newRocket

	case enum.RocketSpeedIncreased:
		msg, err := m.HandleSpeedIncreaseMessage(data)
		if err != nil {
			return entities.Rocket{}, err
		}
		ch := make(chan *entities.Rocket)
		errCh := make(chan error)
		go m.Rockets.GetByID(metadata.Channel, ch, errCh)

		if <-errCh != nil {
			return entities.Rocket{}, err
		}
		rocket.Speed.Update(msg.By)

		ch = make(chan *entities.Rocket)
		errCh = make(chan error)
		go m.Rockets.Update(rocket, ch, errCh)

		if <-errCh != nil {
			return entities.Rocket{}, err
		}

	case enum.RocketSpeedDecreased:
		msg, err := m.HandleSpeedDecreaseMessage(data)
		if err != nil {
			return entities.Rocket{}, err
		}

		ch := make(chan *entities.Rocket)
		errCh := make(chan error)

		go m.Rockets.GetByID(metadata.Channel, ch, errCh)
		if <-errCh != nil {
			return entities.Rocket{}, err
		}

		ch = make(chan *entities.Rocket)
		errCh = make(chan error)

		rocket.Speed.Update(msg.By * -1)
		go m.Rockets.Update(rocket, ch, errCh)

		if <-errCh != nil {
			return entities.Rocket{}, err
		}

	case enum.RocketMissionChanged:
		msg, err := m.HandleMissionChangedMessage(data)
		if err != nil {
			return entities.Rocket{}, err
		}

		ch := make(chan *entities.Rocket)
		errCh := make(chan error)

		go m.Rockets.GetByID(metadata.Channel, ch, errCh)
		if <-errCh != nil {
			return entities.Rocket{}, err
		}

		rocket.Mission = msg.NewMission

		ch = make(chan *entities.Rocket)
		errCh = make(chan error)

		go m.Rockets.Update(rocket, ch, errCh)
		if <-errCh != nil {
			return entities.Rocket{}, err
		}

	case enum.RocketExploded:
		msg, err := m.HandleRocketExplodedMessage(data)
		if err != nil {
			return entities.Rocket{}, err
		}

		ch := make(chan *entities.Rocket)
		errCh := make(chan error)

		go m.Rockets.GetByID(metadata.Channel, ch, errCh)

		if <-errCh != nil {
			return entities.Rocket{}, err
		}

		rocket.Status = entities.RocketStatus{
			Active: false,
			Reason: msg.Reason,
		}

		ch = make(chan *entities.Rocket)
		errCh = make(chan error)

		go m.Rockets.Update(rocket, ch, errCh)

		if <-errCh != nil {
			return entities.Rocket{}, err
		}
	}

	return rocket, nil
}

func NewMessageService() interfaces.MessageService {
	return &MessageService{
		Rockets: entities.NewRocketRepository(),
	}
}
