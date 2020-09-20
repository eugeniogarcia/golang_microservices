package amqp

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	amqphelper "lib/helper/amqp"
	"lib/msgqueue"

	"github.com/streadway/amqp"
)

//Emisor de eventos de amqp. Tenemos una conexión, el exchange por donde se emiten, y un canal de eventos
type amqpEventEmitter struct {
	connection *amqp.Connection
	exchange   string
	//El canal de eventos
	events chan *emittedEvent
}

//El payload y un canal por el que recibir errores
type emittedEvent struct {
	event     msgqueue.Event
	errorChan chan error
}

//NewAMQPEventEmitterFromEnvironment crea un emisor de eventos. Toma los datos para configurar el emisor, los datos de conexión y el nombre del exchange, de variables de entorno
//
//   - AMQP_URL; the URL of the AMQP broker to connect to
//   - AMQP_EXCHANGE; the name of the exchange to bind to
func NewAMQPEventEmitterFromEnvironment() (msgqueue.EventEmitter, error) {
	var url string
	var exchange string

	//Toma los parametros de variables de entorno
	if url = os.Getenv("AMQP_URL"); url == "" {
		url = "amqp://localhost:5672"
	}

	if exchange = os.Getenv("AMQP_EXCHANGE"); exchange == "" {
		exchange = "example"
	}

	//Crea una conexión
	conn := <-amqphelper.RetryConnect(url, 5*time.Second)
	//Crea el emisor de eventos, con la conexión y el exchange
	return NewAMQPEventEmitter(conn, exchange)
}

// NewAMQPEventEmitter creates a new event emitter.
func NewAMQPEventEmitter(conn *amqp.Connection, exchange string) (msgqueue.EventEmitter, error) {
	emitter := amqpEventEmitter{
		connection: conn,
		exchange:   exchange,
	}
	//Crea el exchange
	err := emitter.setup()
	if err != nil {
		return nil, err
	}

	return &emitter, nil
}

//Crea un exchange en un canal
func (a *amqpEventEmitter) setup() error {
	//Obtiene un canal
	channel, err := a.connection.Channel()
	if err != nil {
		return err
	}

	defer channel.Close()

	//Crea un exchange de tipo topic
	//El exchange es persistente, se mantiene incluso cuando se cierra el broker
	err = channel.ExchangeDeclare(a.exchange,
		"topic", //exchange de tipo topic
		true,    //persistent - No se elimina cuando el broker se cierra
		false,   // autodelete - no se elimina cuando el canal se cierre
		false,   //internal - no es interno. Se podrá publicar mensajes
		false,   //noWait - esperamos a que el exchange nos confirme que el mensaje se haya publicado
		nil)
	return err
}

//Publica un evento en el exchange
func (a *amqpEventEmitter) Emit(event msgqueue.Event) error {
	channel, err := a.connection.Channel()
	if err != nil {
		return err
	}

	defer channel.Close()

	// TODO: Alternatives to JSON? Msgpack or Protobuf, maybe?
	jsonBody, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("could not JSON-serialize event: %s", err)
	}

	//Prepara el mensaje. El mensaje tiene un header, content-type y body
	msg := amqp.Publishing{
		Headers:     amqp.Table{"x-event-name": event.EventName()},
		ContentType: "application/json",
		Body:        jsonBody,
	}

	//Publicamos el evento	//routing
	err = channel.Publish(a.exchange,
		event.EventName(), //routing
		false,             //comprobamos que el mesanje haya llegado a una cola por lo menos
		false,             //espera hasta que el mensaje haya sido consumido por al
		msg)               //Mensaje
	return err
}
