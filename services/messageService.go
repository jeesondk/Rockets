package services

import (
	DTO "RocketService/dto"
	"RocketService/entities"
	"RocketService/enum"
	"fmt"
	"sync"
)

type MessageServiceInterface interface {
	HandleSpeedIncreaseMessage(data interface{}) (DTO.RocketSpeedIncreased, error)
	HandleSpeedDecreaseMessage(data interface{}) (DTO.RocketSpeedDecreased, error)
	HandleMissionChangedMessage(data interface{}) (DTO.RocketMissionChanged, error)
	HandleRocketExplodedMessage(data interface{}) (DTO.RocketExploded, error)
	HandleMessage(metadata DTO.MetaData, data interface{}) error
}

type MessageService struct {
	RocketRepository *entities.RocketRepository
	Mux              *sync.Mutex
	Wg               *sync.WaitGroup
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

func (m *MessageService) HandleMessage(metadata DTO.MetaData, data interface{}) error {
	switch metadata.MessageType {
	case enum.RocketLaunched:
		msg, err := m.handleLaunchMessage(data)
		if err != nil {
			return err
		}

		newRocket := entities.Rocket{
			ID:         metadata.Channel,
			RocketType: msg.Type,
			Mission:    msg.Mission,
			Speed: entities.RocketSpeed{
				Current: msg.LaunchSpeed,
				Max:     msg.LaunchSpeed,
			},
			LaunchDate: metadata.MessageTime,
			Status: entities.RocketStatus{
				Active: true,
				Reason: "",
			},
			EventCursor: 1,
		}

		if newRocket.ID != "" {
			m.Wg.Add(1)
			m.RocketRepository.Create(&newRocket, m.Mux, m.Wg)
			m.Wg.Wait()
			return nil
		}
		return fmt.Errorf("unable to create rocket, missing ID")

	default:
		m.Wg.Add(1)
		err := m.RocketRepository.AddEvent(&metadata, data, m.Mux, m.Wg)
		m.Wg.Wait()
		if err != nil {
			return err
		}
		return nil
	}
}

func (m *MessageService) handleLaunchMessage(data interface{}) (DTO.RocketLaunched, error) {
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
func NewMessageService() MessageService {
	return MessageService{}
}
