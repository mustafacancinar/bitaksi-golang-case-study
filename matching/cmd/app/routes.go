package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {
	router := mux.NewRouter()
	router.Handle("/matching", JWTMiddleware(http.HandlerFunc(MatchingHandler))).Methods("POST")
	router.HandleFunc("/token", GenerateTokenHandler).Methods("POST")
	return router
}