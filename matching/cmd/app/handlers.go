package app

import (
	"encoding/json"
	"net/http"

	"github.com/cinarizasyon/bitaksi-golang-case-study/matching/internal"
	"github.com/go-playground/validator/v10"
)

func MatchingHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var request internal.MatchingRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if handleValidationError(w, request) != nil {
		return
	}


	service := internal.NewMatchingService("http://localhost:8080")
	
	// TODO: refactor this logic with a gorutine that checks the health of the remote service
	if isHealthy, _:= service.CheckRemoteServiceHealth(); !isHealthy{
		// TODO: log error
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	response, err := service.Match(&request)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if response == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func GenerateTokenHandler(w http.ResponseWriter, r *http.Request) {
	token, err := internal.GenerateJWT()
	if err != nil {
		http.Error(w, "Something went wrong. Please try again later.", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
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