package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Location struct {
	Type        string    `bson:"type"`
	Coordinates []float64 `bson:"coordinates"`
}

type DriverLocation struct {
	ID 	  primitive.ObjectID `bson:"_id,omitempty"`
	Location Location		  `json:"location"`
}