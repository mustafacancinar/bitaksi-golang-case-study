package app

import (
	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/healthz", HealthCheckHandler).Methods("GET")
	router.HandleFunc("/drivers", CreateHandler).Methods("POST")
	router.HandleFunc("/drivers/bulk", BulkCreateHandler).Methods("POST")
	router.HandleFunc("/drivers/search", SearchHandler).Methods("POST")
	return router
}