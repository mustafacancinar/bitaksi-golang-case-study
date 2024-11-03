package models

type CreateDriverLocationRequest struct {
	Longitude float64 `json:"longitude" validate:"required,lte=180,gte=-180"`
	Latitude  float64 `json:"latitude" validate:"required,lte=90,gte=-90"`
}

type BulkCreateDriverLocationRequest struct {
	Locations []CreateDriverLocationRequest `json:"locations" validate:"required,dive"`
}

type SearchDriverLocationRequest struct {
	Longitude float64 `json:"longitude" validate:"required,lte=180,gte=-180"`
	Latitude  float64 `json:"latitude" validate:"required,lte=90,gte=-90"`
	Radius    float64 `json:"radius" validate:"required,gte=0"`
}