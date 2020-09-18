package configuration

import (
	"encoding/json"
	"fmt"
	"os"

	"lib/persistence/dblayer"
)

//Define constantes
var (
	DBTypeDefault       = dblayer.DBTYPE("mongodb")
	DBConnectionDefault = "mongodb://127.0.0.1"
	RestfulEPDefault    = "localhost:8181"
)

//Define el tipo que equivale al json de configuracion
type ServiceConfig struct {
	Databasetype    dblayer.DBTYPE `json:"databasetype"`
	DBConnection    string         `json:"dbconnection"`
	RestfulEndpoint string         `json:"restfulapi_endpoint"`
}

//ExtractConfiguration obtiene la configuración desde un archivo de configuración
func ExtractConfiguration(filename string) (ServiceConfig, error) {
	//Configuracion por defecto
	conf := ServiceConfig{
		DBTypeDefault,
		DBConnectionDefault,
		RestfulEPDefault,
	}

	//Lee el archivo
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Configuration file not found. Continuing with default values.")
		return conf, err
	}

	//Serializa el contenido del archivo en nuestro tipo
	err = json.NewDecoder(file).Decode(&conf)
	return conf, err
}
