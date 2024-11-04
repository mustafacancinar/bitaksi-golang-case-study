
package internal

type MatchingRequest struct {
	Longitude float64 `json:"longitude" validate:"required,lte=180,gte=-180"`
	Latitude  float64 `json:"latitude" validate:"required,lte=90,gte=-90"`
	Radius    float64 `json:"radius" validate:"required,gte=0"`
}

type MatchingResponse struct {
	DriverId string `json:"id"`
	Distance float64 `json:"distance"`
	Longitude float64 `json:"longitude"`
	Latitude float64 `json:"latitude"`
}