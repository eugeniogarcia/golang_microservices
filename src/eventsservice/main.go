package main

import (
	"flag"
	"log"

	"eventsservice/rest"

	"lib/configuration"
	"lib/persistence/dblayer"
)

func main() {

	//Toma del flag conf la ruta del archivo de configuracion
	confPath := flag.String("conf", `.\configuration\config.json`, "flag to set the path to the configuration json file")

	//Parsea los argumentos
	flag.Parse()

	//Lee la configuracion
	config, _ := configuration.ExtractConfiguration(*confPath)

	log.Println("Connecting to database")
	dbhandler, err := dblayer.NewPersistenceLayer(config.Databasetype, config.DBConnection)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database connection successful... ")
	//RESTful API start
	httpErrChan, httptlsErrChan := rest.ServeAPI(config.RestfulEndpoint, config.RestfulTLSEndPint, dbhandler)
	select {
	case err := <-httpErrChan:
		log.Fatal("HTTP Error: ", err)
	case err := <-httptlsErrChan:
		log.Fatal("HTTPS Error: ", err)
	}
}
