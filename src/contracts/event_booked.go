package contracts

//Implementa la interface event porque implementa el metodo EventName()
// EventBookedEvent is emitted whenever an event is booked
type EventBookedEvent struct {
	EventID string `json:"eventId"`
	UserID  string `json:"userId"`
}

// EventName returns the event's name
func (c *EventBookedEvent) EventName() string {
	return "eventBooked"
}
