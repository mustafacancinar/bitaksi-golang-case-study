package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/swaggo/http-swagger"
	_ "github.com/cinarizasyon/bitaksi-golang-case-study/matching/cmd/docs"
)

func RegisterRoutes() *mux.Router {
	router := mux.NewRouter()
	router.Handle("/matching", JWTMiddleware(http.HandlerFunc(MatchingHandler))).Methods("POST")
	router.HandleFunc("/token", GenerateTokenHandler).Methods("POST")
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	return router
}