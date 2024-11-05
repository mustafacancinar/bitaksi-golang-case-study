package app

import (
	"net/http"

	_ "github.com/cinarizasyon/bitaksi-golang-case-study/driver-location/cmd/docs"
	"github.com/gorilla/mux"
	"github.com/swaggo/http-swagger"
)

func RegisterRoutes() *mux.Router {
	router := mux.NewRouter()
	router.Handle("/healthz", JWTMiddleware(http.HandlerFunc(HealthCheckHandler))).Methods("GET")
	router.Handle("/drivers", JWTMiddleware(http.HandlerFunc(CreateHandler))).Methods("POST")
	router.Handle("/drivers/bulk", JWTMiddleware(http.HandlerFunc(BulkCreateHandler))).Methods("POST")
	router.Handle("/drivers/search", JWTMiddleware(http.HandlerFunc(SearchHandler))).Methods("POST")
	router.Handle("/drivers/upload", JWTMiddleware(http.HandlerFunc(UploadHandler))).Methods("POST")
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	return router
}