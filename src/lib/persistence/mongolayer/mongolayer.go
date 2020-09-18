package mongolayer

import (
	"lib/persistence"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//Constantes
const (
	DB     = "myevents"
	USERS  = "users"
	EVENTS = "events"
)

//Tipo custom. Contiene una sesión de mongo
type MongoDBLayer struct {
	session *mgo.Session
}

//Constructor
func NewMongoDBLayer(connection string) (persistence.DatabaseHandler, error) {
	//Nos conectamos
	s, err := mgo.Dial(connection)
	//retorna el objeto
	return &MongoDBLayer{
		session: s,
	}, err
}

//Define los métodos que gestionan la capa de persistencia
//Añade un evento
func (mgoLayer *MongoDBLayer) AddEvent(e persistence.Event) ([]byte, error) {

	//toma una sesión
	s := mgoLayer.getFreshSession()
	defer s.Close()

	//Valida los datos
	if !e.ID.Valid() {
		e.ID = bson.NewObjectId()
	}

	if !e.Location.ID.Valid() {
		e.Location.ID = bson.NewObjectId()
	}

	//Inserta en la colección de una base de datos
	return []byte(e.ID), s.DB(DB).C(EVENTS).Insert(e)
}

func (mgoLayer *MongoDBLayer) FindEvent(id []byte) (persistence.Event, error) {
	s := mgoLayer.getFreshSession()
	defer s.Close()
	e := persistence.Event{}

	//Convierte el slice de bytes en un bson
	err := s.DB(DB).C(EVENTS).FindId(bson.ObjectId(id)).One(&e)
	return e, err
}

func (mgoLayer *MongoDBLayer) FindEventByName(name string) (persistence.Event, error) {
	s := mgoLayer.getFreshSession()
	defer s.Close()
	e := persistence.Event{}

	//Busca en una colección documentos que contengan un campo
	err := s.DB(DB).C(EVENTS).Find(bson.M{"name": name}).One(&e)
	return e, err
}

func (mgoLayer *MongoDBLayer) FindAllAvailableEvents() ([]persistence.Event, error) {
	s := mgoLayer.getFreshSession()
	defer s.Close()
	events := []persistence.Event{}
	//Busca sin especificar criterios, y retorna todo lo que encuentra en un slice
	err := s.DB(DB).C(EVENTS).Find(nil).All(&events)
	return events, err
}

func (mgoLayer *MongoDBLayer) getFreshSession() *mgo.Session {
	return mgoLayer.session.Copy()
}
