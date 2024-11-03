package app

import (
	"net/http"

	"github.com/cinarizasyon/bitaksi-golang-case-study/driver-location/internal"
)

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from create handler"))
}

func BulkCreateHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from bulk create handler"))
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from search handler"))
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	dbClient, err := internal.OpenConnection()
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}
	defer internal.CloseConnection(dbClient)

	err = internal.Ping(dbClient)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	w.WriteHeader(http.StatusOK)
}
