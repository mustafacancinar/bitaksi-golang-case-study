package app

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/cinarizasyon/bitaksi-golang-case-study/matching/internal"
	"github.com/go-playground/validator/v10"
)

// @Summary Matching
// @Description Match a driver with a passenger
// @Tags matching
// @Accept json
// @Produce json
// @Param request body internal.MatchingRequest true "Matching request"
// @Success 200 {object} internal.MatchingResponse
// @Failure 400 {string} string "Invalid JSON"
// @Failure 404 {string} string "No driver found"
// @Failure 500 {string} string "Internal server error"
// @Router /matching [post]
// @Security ApiKeyAuth
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


	service := internal.NewMatchingService(os.Getenv("DRIVER_LOCATION_SERVICE_URL"))
	
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

// @Summary Generate token
// @Description Generate a JWT token
// @Tags token
// @Produce json
// @Success 200 {object} map[string]string
// @Failure 500 {string} string "Internal server error"
// @Router /token [post]
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