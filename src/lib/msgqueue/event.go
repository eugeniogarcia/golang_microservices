package msgqueue

//Event un evento. Con EventName obtenemos el binding, y lo usaremos para saber el tipo de evento
type Event interface {
	EventName() string
}
