package amqp

import (
	"fmt"
	"os"
	"time"

	amqphelper "lib/helper/amqp"
	"lib/msgqueue"

	"github.com/streadway/amqp"
)

const eventNameHeader = "x-event-name"

type amqpEventListener struct {
	connection *amqp.Connection
	exchange   string
	queue      string
	mapper     msgqueue.EventMapper
}

// NewAMQPEventListenerFromEnvironment crea un listener para rabbit. Toma la configuración del entorno
//   - AMQP_URL; the URL of the AMQP broker to connect to
//   - AMQP_EXCHANGE; the name of the exchange to bind to
//   - AMQP_QUEUE; the name of the queue to bind and subscribe
func NewAMQPEventListenerFromEnvironment() (msgqueue.EventListener, error) {
	var url string
	var exchange string
	var queue string

	if url = os.Getenv("AMQP_URL"); url == "" {
		url = "amqp://localhost:5672"
	}

	if exchange = os.Getenv("AMQP_EXCHANGE"); exchange == "" {
		exchange = "example"
	}

	if queue = os.Getenv("AMQP_QUEUE"); queue == "" {
		queue = "example"
	}

	//Crea la conexión
	conn := <-amqphelper.RetryConnect(url, 5*time.Second)
	//Retorna el listener
	return NewAMQPEventListener(conn, exchange, queue)
}

// NewAMQPEventListener crea un listener
// Crea un canal nuevo con cada listener. Los canales no son threadsafe, asíque creamos uno por cliente
func NewAMQPEventListener(conn *amqp.Connection, exchange string, queue string) (msgqueue.EventListener, error) {
	listener := amqpEventListener{
		connection: conn,
		exchange:   exchange,
		queue:      queue,
		mapper:     msgqueue.NewEventMapper(),
	}

	//Crea la cola, y el exchange
	err := listener.setup()
	if err != nil {
		return nil, err
	}

	return &listener, nil
}

func (a *amqpEventListener) setup() error {
	//Crea el canal
	channel, err := a.connection.Channel()
	if err != nil {
		return err
	}

	//Lo cierra al terminar
	defer channel.Close()

	//Crea el exchange
	err = channel.ExchangeDeclare(a.exchange, "topic", true, false, false, false, nil)
	if err != nil {
		return err
	}

	//Crea la cola
	_, err = channel.QueueDeclare(a.queue, true, false, false, false, nil)
	if err != nil {
		return fmt.Errorf("could not declare queue %s: %s", a.queue, err)
	}

	return nil
}

// Listen configures the event listener to listen for a set of events that are
// specified by name as parameter.
// This method will return two channels: One will contain successfully decoded
// events, the other will contain errors for messages that could not be
// successfully decoded.

//Listen escucha una serie de eventos, y retorna el canal por el que se emitiran los mensajes
func (l *amqpEventListener) Listen(eventNames ...string) (<-chan msgqueue.Event, <-chan error, error) {
	//Obtiene un canal
	channel, err := l.connection.Channel()
	if err != nil {
		return nil, nil, err
	}

	//Para cada evento vinculamos nuestra cola al exchange. De este modo por nuestra cola iran llegando los eventos que publique el exchange que sean de nuestro interes
	for _, event := range eventNames {
		//Cre un binding entre la cola y el exchange, usando como binding el nombre del evento. Los argumentos son el nombre de la cola, el binding, el nombre del exchange y si hacemos acknowledge automatico - no lo hacemos
		if err := channel.QueueBind(l.queue, event, l.exchange, false, nil); err != nil {
			return nil, nil, fmt.Errorf("could not bind event %s to queue %s: %s", event, l.queue, err)
		}
	}

	//Consume de una cola. Obtiene un canal
	msgs, err := channel.Consume(l.queue, //Nombre de la cola
		"",    //el identificador del cliente. Si no indicamos nada, se autogenera uno
		false, //No hacemos auto acknowledge
		false, //No es exclusivo, así que puede haber varios consumidores en la misma cola
		false,
		false,
		nil)
	if err != nil {
		return nil, nil, fmt.Errorf("could not consume queue: %s", err)
	}

	events := make(chan msgqueue.Event)
	errors := make(chan error)

	//Procesa los eventos que lleguen por el canal, hasta que el canal se cierre, en una go rutina
	go func() {
		//Un stream de mensajes
		for msg := range msgs {
			//En el mensaje, en la cabecera especificamos el tipo de evento
			rawEventName, ok := msg.Headers[eventNameHeader]
			if !ok {
				errors <- fmt.Errorf("message did not contain %s header", eventNameHeader)
				msg.Nack(false, false)
				continue
			}
			//Obtiene el nombre del evento
			eventName, ok := rawEventName.(string)
			if !ok {
				errors <- fmt.Errorf("header %s did not contain string", eventNameHeader)
				//Informa que no se proceso el mensaje, para que vuelva a la cola
				msg.Nack(false, false)
				continue
			}

			//Mapea los datos a nuestro tipo de evento
			event, err := l.mapper.MapEvent(eventName, msg.Body)
			if err != nil {
				errors <- fmt.Errorf("could not unmarshal event %s: %s", eventName, err)
				//Informa que no se proceso el mensaje, para que vuelva a la cola
				msg.Nack(false, false)
				continue
			}

			//Envia el evento al canal
			events <- event
			//Hace el acknowledgment
			msg.Ack(false)
		}
	}()

	return events, errors, nil
}

//Mapper el mapa
func (l *amqpEventListener) Mapper() msgqueue.EventMapper {
	return l.mapper
}
