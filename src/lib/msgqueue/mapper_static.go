package msgqueue

import (
	"encoding/json"
	"fmt"

	"contracts"

	"github.com/mitchellh/mapstructure"
)

type StaticEventMapper struct{}

//MapEvent retorna una interface Event
func (e *StaticEventMapper) MapEvent(eventName string, serialized interface{}) (Event, error) {
	var event Event

	//Determina el tipo de evento, y retorna el tipo correspondiente
	switch eventName {
	case "eventCreated":
		event = &contracts.EventCreatedEvent{}
	case "locationCreated":
		event = &contracts.LocationCreatedEvent{}
	case "eventBooked":
		event = &contracts.EventBookedEvent{}
	default:
		return nil, fmt.Errorf("unknown event type %s", eventName)
	}

	switch s := serialized.(type) {
	case []byte:
		err := json.Unmarshal(s, event)
		if err != nil {
			return nil, fmt.Errorf("could not unmarshal event %s: %s", eventName, err)
		}
	default:
		cfg := mapstructure.DecoderConfig{
			Result:  event,
			TagName: "json",
		}
		dec, err := mapstructure.NewDecoder(&cfg)
		if err != nil {
			return nil, fmt.Errorf("could not initialize decoder for event %s: %s", eventName, err)
		}

		err = dec.Decode(s)
		if err != nil {
			return nil, fmt.Errorf("could not unmarshal event %s: %s", eventName, err)
		}
	}

	return event, nil
}