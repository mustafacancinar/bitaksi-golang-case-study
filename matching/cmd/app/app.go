package app

import (
	"net/http"
)

func Run() {
	router := RegisterRoutes()
	http.ListenAndServe(":8081", router)
}