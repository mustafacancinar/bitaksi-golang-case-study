package app

import (
	"net/http"
)

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from create handler"))
}

func BulkCreateHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from bulk create handler"))
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from search handler"))
}
