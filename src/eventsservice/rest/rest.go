package rest

import (
	"net/http"

	"lib/persistence"

	"github.com/gorilla/mux"
)

func ServeAPI(endpoint string, databasehandler persistence.DatabaseHandler) error {
	//Crea el handler
	handler := NewEventHandler(databasehandler)
	//Crea un router
	r := mux.NewRouter()

	//Define las rutas
	eventsrouter := r.PathPrefix("/events").Subrouter()
	eventsrouter.Methods("GET").Path("/{SearchCriteria}/{search}").HandlerFunc(handler.FindEventHandler)
	eventsrouter.Methods("GET").Path("").HandlerFunc(handler.AllEventHandler)
	eventsrouter.Methods("POST").Path("").HandlerFunc(handler.NewEventHandler)

	//Empieza a escuchar
	return http.ListenAndServe(endpoint, r)
}
