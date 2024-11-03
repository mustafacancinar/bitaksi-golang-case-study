package app

type CreateDriverLocationRequest struct {
	Longitude float64 `json:"longitude" validate:"lte=180,gte=-180"`
	Latitude  float64 `json:"latitude" validate:"lte=90,gte=-90"`
}
