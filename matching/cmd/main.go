package main

import (
	"net/http"
)

func driversHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from matching handler"))
}

func main() {
	http.HandleFunc("/drivers", driversHandler)
	http.ListenAndServe(":8081", nil)
}