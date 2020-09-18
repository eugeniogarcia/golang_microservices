package dblayer

import (
	"lib/persistence"
	"lib/persistence/mongolayer"
)

//DBTYPE define un tipo custom
type DBTYPE string

//define constantes
const (
	MONGODB    DBTYPE = "mongodb"
	DOCUMENTDB DBTYPE = "documentdb"
	DYNAMODB   DBTYPE = "dynamodb"
)

//NewPersistenceLayer Devuelve un interface con los metodos necesarios para administrar los datos en la capa de persistencia
func NewPersistenceLayer(options DBTYPE, connection string) (persistence.DatabaseHandler, error) {

	switch options {
	case MONGODB:
		return mongolayer.NewMongoDBLayer(connection)
	}
	return nil, nil
}
