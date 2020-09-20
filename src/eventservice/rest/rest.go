package rest

import (
	"net/http"

	"lib/msgqueue"
	"lib/persistence"

	"github.com/gorilla/mux"
)

//ServeAPI arranca el servidor http/https
func ServeAPI(endpoint string, tlsendpoint string, dbHandler persistence.DatabaseHandler, eventEmitter msgqueue.EventEmitter) (chan error, chan error) {
	handler := newEventHandler(dbHandler, eventEmitter)

	r := mux.NewRouter()
	eventsrouter := r.PathPrefix("/events").Subrouter()
	eventsrouter.Methods("GET").Path("/{SearchCriteria}/{search}").HandlerFunc(handler.findEventHandler)
	eventsrouter.Methods("GET").Path("").HandlerFunc(handler.allEventHandler)
	eventsrouter.Methods("GET").Path("/{eventID}").HandlerFunc(handler.oneEventHandler)
	eventsrouter.Methods("POST").Path("").HandlerFunc(handler.newEventHandler)

	locationRouter := r.PathPrefix("/locations").Subrouter()
	locationRouter.Methods("GET").Path("").HandlerFunc(handler.allLocationsHandler)
	locationRouter.Methods("POST").Path("").HandlerFunc(handler.newLocationHandler)

	httpErrChan := make(chan error)
	httptlsErrChan := make(chan error)

	go func() {
		httptlsErrChan <- http.ListenAndServeTLS(tlsendpoint, "cert.pem", "key.pem", r)
	}()

	go func() {
		httpErrChan <- http.ListenAndServe(endpoint, r)
	}()

	return httpErrChan, httptlsErrChan
}
