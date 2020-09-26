package rest

import (
	"net/http"
	"time"

	"lib/msgqueue"
	"lib/persistence"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func ServeAPI(listenAddr string, database persistence.DatabaseHandler, eventEmitter msgqueue.EventEmitter) error {
	r := mux.NewRouter()
	usuario, err := database.AddUser(persistence.User{First: "Eugenio", Last: "Garcia", Age: 50})
	if err != nil {
		return err
	}
	r.Methods("post").Path("/events/{eventID}/bookings").Handler(&CreateBookingHandler{eventEmitter, database, usuario})

	srv := http.Server{
		Handler:      handlers.CORS()(r),
		Addr:         listenAddr,
		WriteTimeout: 2 * time.Second,
		ReadTimeout:  1 * time.Second,
	}

	return srv.ListenAndServe()
}
