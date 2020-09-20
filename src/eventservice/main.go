package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"net/http"

	"github.com/Shopify/sarama"

	"eventservice/rest"
	"lib/configuration"
	"lib/msgqueue"
	msgqueue_amqp "lib/msgqueue/amqp"
	"lib/msgqueue/kafka"
	"lib/persistence/dblayer"

	"github.com/streadway/amqp"
)

func main() {
	var eventEmitter msgqueue.EventEmitter

	confPath := flag.String("conf", `.\configuration\config.json`, "flag to set the path to the configuration json file")
	flag.Parse()
	//extract configuration
	config, _ := configuration.ExtractConfiguration(*confPath)

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

		if err != nil {
			panic(err)
		}

		eventEmitter, err = msgqueue_amqp.NewAMQPEventEmitter(conn, "events")
		if err != nil {
			panic(err)
		}
	case "kafka":
		conf := sarama.NewConfig()
		conf.Producer.Return.Successes = true
		conn, err := sarama.NewClient(config.KafkaMessageBrokers, conf)
		if err != nil {
			panic(err)
		}

		eventEmitter, err = kafka.NewKafkaEventEmitter(conn)
		if err != nil {
			panic(err)
		}
	default:
		panic("Bad message broker type: " + config.MessageBrokerType)
	}

	fmt.Println("Connecting to database")
	dbhandler, _ := dblayer.NewPersistenceLayer(config.Databasetype, config.DBConnection)

	go func() {
		fmt.Println("Serving metrics API")
		h := http.NewServeMux()
		h.Handle("/metrics", promhttp.Handler())

		http.ListenAndServe(":9100", h)
	}()

	fmt.Println("Serving API")
	//RESTful API start
	httpErrChan, httptlsErrChan := rest.ServeAPI(config.RestfulEndpoint, config.RestfulTLSEndPint, dbhandler, eventEmitter)

	select {
	case err := <-httpErrChan:
		log.Fatal("Http error: ", err)
	case err := <-httptlsErrChan:
		log.Fatal("Https error: ", err)
	}
}
