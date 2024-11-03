package app

import (
	"github.com/cinarizasyon/bitaksi-golang-case-study/driver-location/internal"
	"net/http"
)

func Run() {
	internal.InitDatabase()
	router := RegisterRoutes()
	http.ListenAndServe(":8080", router)
}