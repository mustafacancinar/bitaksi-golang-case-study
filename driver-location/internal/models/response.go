package models

type DriverLocationSearchResponse struct {
	Id 	  any    `json:"id"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
	Distance float64 `json:"distance"`
}