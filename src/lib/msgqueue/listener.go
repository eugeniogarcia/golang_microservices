package msgqueue

// EventListener describes an interface for a class that can listen to events. Un metodo que devuelve un canal por el que se recibiran los mensajes. Otro método que sirve para mapear los datos
type EventListener interface {
	Listen(events ...string) (<-chan Event, <-chan error, error)
	Mapper() EventMapper
}
