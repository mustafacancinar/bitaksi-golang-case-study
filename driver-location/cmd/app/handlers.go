package app

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"

	"github.com/cinarizasyon/bitaksi-golang-case-study/driver-location/internal"
	"github.com/cinarizasyon/bitaksi-golang-case-study/driver-location/internal/models"
)

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var request models.CreateDriverLocationRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	err = validate.Struct(request)
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		errors := FormatValidationErrors(validationErrors)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errors)
		return
	}

	service := internal.NewLocationDriverService()
	defer internal.CloseConnection(service.Client)

	err = service.CreateDriverLocation(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func BulkCreateHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var request models.BulkCreateDriverLocationRequest

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	err = validate.Struct(request)
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		errors := FormatValidationErrors(validationErrors)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errors)
		return
	}

	service := internal.NewLocationDriverService()
	defer internal.CloseConnection(service.Client)

	err = service.BulkCreateDriverLocations(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from search handler"))
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	dbClient, err := internal.OpenConnection()
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}
	defer internal.CloseConnection(dbClient)

	err = internal.Ping(dbClient)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	w.WriteHeader(http.StatusOK)
}
