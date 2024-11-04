package app

import (
	"fmt"
	"net/http"
	"os"
)

func Run() {
	router := RegisterRoutes()
	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), router)
}