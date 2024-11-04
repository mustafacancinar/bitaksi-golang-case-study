package app

import (
	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/matching", MatchingHandler).Methods("POST")
	return router
}