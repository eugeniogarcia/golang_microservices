package main

import (
	"flag"
	"log"
	"time"

	"bookingservice/listener"
	"bookingservice/rest"
	"lib/configuration"
	"lib/msgqueue"
	msgqueue_amqp "lib/msgqueue/amqp"
	"lib/msgqueue/kafka"
	"lib/persistence/dblayer"

	"github.com/Shopify/sarama"
	"github.com/streadway/amqp"
)

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	//Escuchamos ...
	var eventListener msgqueue.EventListener
	//... y emitimos mensajes
	var eventEmitter msgqueue.EventEmitter

	confPath := flag.String("conf", "./configuration/config.json", "flag to set the path to the configuration json file")
	flag.Parse()

	//extract configuration
	config, _ := configuration.ExtractConfiguration(*confPath)

	//Comprobamos que tipo de broker queremos usar
	switch config.MessageBrokerType {
	case "amqp":

		var conn *amqp.Connection
		var err error
		for i := 0; i < 3; i++ {
			conn, err = amqp.Dial(config.AMQPMessageBroker)
			if err == nil {
				log.Println("connection successfully established")
				break
			}

			log.Printf("AMQP connection failed with error: %s", err)
			time.Sleep(5 * time.Second)
		}

		panicIfErr(err)

		//Escuchamos por eventos events, y booking
		eventListener, err = msgqueue_amqp.NewAMQPEventListener(conn, "events", "booking")
		panicIfErr(err)

		//A su vez publicamos events
		eventEmitter, err = msgqueue_amqp.NewAMQPEventEmitter(conn, "events")
		panicIfErr(err)
	case "kafka":
		conf := sarama.NewConfig()
		conf.Producer.Return.Successes = true
		conn, err := sarama.NewClient(config.KafkaMessageBrokers, conf)
		panicIfErr(err)

		eventListener, err = kafka.NewKafkaEventListener(conn, []int32{})
		panicIfErr(err)

		eventEmitter, err = kafka.NewKafkaEventEmitter(conn)
		panicIfErr(err)
	default:
		panic("Bad message broker type: " + config.MessageBrokerType)
	}

	dbhandler, _ := dblayer.NewPersistenceLayer(config.Databasetype, config.DBConnection)

	//Procesa los eventos en una go rutina
	processor := listener.EventProcessor{eventListener, dbhandler}
	go processor.ProcessEvents()

	//También sirve peticiones http. Usa el emiter para publicar los cambios
	panicIfErr(rest.ServeAPI(config.RestfulEndpoint, dbhandler, eventEmitter))
}
