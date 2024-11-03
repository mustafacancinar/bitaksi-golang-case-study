package models

type Location struct {
	Type        string    `bson:"type"`
	Coordinates []float64 `bson:"coordinates"`
}

type DriverLocation struct {
	ID 	  any `bson:"_id,omitempty"`
	Location Location		  `json:"location"`
}