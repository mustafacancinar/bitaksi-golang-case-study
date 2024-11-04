package app

import (
	"fmt"
	"net/http"
	"os"

	"github.com/cinarizasyon/bitaksi-golang-case-study/driver-location/internal"
)

func Run() {
	internal.InitDatabase()
	router := RegisterRoutes()
	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), router)
}