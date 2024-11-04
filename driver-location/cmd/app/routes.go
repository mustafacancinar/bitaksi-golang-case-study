package app

import (
	"github.com/gorilla/mux"
	"github.com/swaggo/http-swagger"
	_ "github.com/cinarizasyon/bitaksi-golang-case-study/driver-location/cmd/docs"
)

func RegisterRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/healthz", HealthCheckHandler).Methods("GET")
	router.HandleFunc("/drivers", CreateHandler).Methods("POST")
	router.HandleFunc("/drivers/bulk", BulkCreateHandler).Methods("POST")
	router.HandleFunc("/drivers/search", SearchHandler).Methods("POST")
	router.HandleFunc("/drivers/upload", UploadHandler).Methods("POST")
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	return router
}