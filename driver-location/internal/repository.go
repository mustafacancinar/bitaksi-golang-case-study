package internal

import (
	"fmt"
	"context"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Repository interface {
	CreateDriverLocation(ctx context.Context, location *DriverLocation) error
	BulkCreateDriverLocations(ctx context.Context, locations []*DriverLocation) error
	SearchLocation(ctx context.Context, longitude, latitude float64, radius int) ([]DriverLocation, error)
}

type DriverLocationRepository struct {
	collection *mongo.Collection
}

func NewDriverLocationRepository(db *mongo.Database, collectionName string) *DriverLocationRepository {
	return &DriverLocationRepository{collection: db.Collection(collectionName)}
}

func (r *DriverLocationRepository) CreateDriverLocation(ctx context.Context, location *DriverLocation) error {
	_, err := r.collection.InsertOne(ctx, location)
	if err != nil {
		return fmt.Errorf("failed to insert driver location: %w", err)
	}

	return nil
}

func (r *DriverLocationRepository) BulkCreateDriverLocations(ctx context.Context, locations []*DriverLocation) error {		
	_, err := r.collection.InsertMany(ctx, locations)
	if err != nil {
		return fmt.Errorf("failed to insert driver locations: %w", err)
	}

	return nil
}

func (r *DriverLocationRepository) SearchLocation(ctx context.Context, longitude, latitude float64, radius int) ([]DriverLocation, error) {
	filter := bson.M{
		"location": bson.M{
			"$geoWithin": bson.M{
				"$centerSphere": []any{[]float64{longitude, latitude}, float64(radius) / 3963.2},
			},
		},
	}

	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to find driver locations: %w", err)
	}
	defer cursor.Close(ctx)

	var locations []DriverLocation
	if err := cursor.All(ctx, &locations); err != nil {
		return nil, fmt.Errorf("failed to decode driver locations: %w", err)
	}

	return locations, nil
}

