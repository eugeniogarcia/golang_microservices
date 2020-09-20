package contracts

import "lib/persistence"

//Implementa la interface event porque implementa el metodo EventName()
// LocationCreatedEvent is emitted whenever a location is created
type LocationCreatedEvent struct {
	ID      string             `json:"id"`
	Name    string             `json:"name"`
	Address string             `json:"address"`
	Country string             `json:"country"`
	Halls   []persistence.Hall `json:"halls"`
}

// EventName returns the event's name
func (c *LocationCreatedEvent) EventName() string {
	return "locationCreated"
}
