package app

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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

	if handleValidationError(w, request) != nil {
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

	if handleValidationError(w, request) != nil {
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

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Invalid file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	req := models.BulkCreateDriverLocationRequest{}
	locations := make([]models.CreateDriverLocationRequest, 0)
	csvReader := csv.NewReader(file)
	lines, err := csvReader.ReadAll()

	if err != nil {
		http.Error(w, "Invalid CSV", http.StatusBadRequest)
		return
	}

	for i, line := range lines {
		if i == 0 {
			continue
		}

		longitude, err := strconv.ParseFloat(line[0], 64)
		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid longitude: %s", line[0]), http.StatusBadRequest)
			return
		}

		latitude, err := strconv.ParseFloat(line[1], 64)
		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid latitude: %s", line[1]), http.StatusBadRequest)
			return
		}

		locations = append(locations, models.CreateDriverLocationRequest{
			Longitude: longitude,
			Latitude:  latitude,
		})
	}
	req.Locations = locations

	if handleValidationError(w, req) != nil {
		return
	}

	service := internal.NewLocationDriverService()
	defer internal.CloseConnection(service.Client)

	err = service.BulkCreateDriverLocations(&req)
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

func handleValidationError(w http.ResponseWriter, req any) error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(req)
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		errors := FormatValidationErrors(validationErrors)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errors)
		return err
	}

	return nil
}