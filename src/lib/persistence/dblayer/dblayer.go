package dblayer

import (
	"lib/persistence"
	"lib/persistence/mongolayer"
)

//define un tipo custom
type DBTYPE string

//define constantes
const (
	MONGODB  DBTYPE = "mongodb"
	DYNAMODB DBTYPE = "dynamodb"
)

//Devuelve un interface con los metodos necesarios para administrar los datos en la capa de persistencia
func NewPersistenceLayer(options DBTYPE, connection string) (persistence.DatabaseHandler, error) {

	switch options {
	case MONGODB:
		return mongolayer.NewMongoDBLayer(connection)
	}
	return nil, nil
}