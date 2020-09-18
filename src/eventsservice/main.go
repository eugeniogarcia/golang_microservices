package main

import (
	"flag"
	"fmt"
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

	fmt.Println("Connecting to database")
	//Crea la capa de persistencia con el tipo - mongo - y los datos de conexión a la base de datos
	dbhandler, _ := dblayer.NewPersistenceLayer(config.Databasetype, config.DBConnection)

	//RESTful API start
	//arranca el servidor
	log.Fatal(rest.ServeAPI(config.RestfulEndpoint, dbhandler))
}
