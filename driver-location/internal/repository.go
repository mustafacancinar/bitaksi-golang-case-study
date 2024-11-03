package internal

import (
	"context"
	"fmt"

	"github.com/cinarizasyon/bitaksi-golang-case-study/driver-location/internal/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Repository interface {
	CreateDriverLocation(ctx context.Context, location *models.DriverLocation) error
	BulkCreateDriverLocations(ctx context.Context, locations []*models.DriverLocation) error
	SearchLocation(ctx context.Context, longitude, latitude float64, radius int) ([]models.DriverLocation, error)
}

type DriverLocationRepository struct {
	collection *mongo.Collection
}

func NewDriverLocationRepository(db *mongo.Database, collectionName string) *DriverLocationRepository {
	return &DriverLocationRepository{collection: db.Collection(collectionName)}
}

func (r *DriverLocationRepository) CreateDriverLocation(ctx context.Context, location *models.DriverLocation) error {
	_, err := r.collection.InsertOne(ctx, location)
	if err != nil {
		return fmt.Errorf("failed to insert driver location: %w", err)
	}

	return nil
}

func (r *DriverLocationRepository) BulkCreateDriverLocations(ctx context.Context, locations []*models.DriverLocation) error {		
	_, err := r.collection.InsertMany(ctx, locations)
	if err != nil {
		return fmt.Errorf("failed to insert driver locations: %w", err)
	}

	return nil
}

func (r *DriverLocationRepository) SearchLocation(ctx context.Context, longitude, latitude float64, radius int) ([]models.DriverLocation, error) {
	filter := bson.M{
		"location": bson.M{
			"$geoWithin": bson.M{
				"$centerSphere": []any{[]float64{longitude, latitude}, float64(radius) / 6378.1}, // The equatorial radius of the Earth is approximately 3,963.2 miles or 6,378.1 kilometers.
			},
		},
	}

	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to find driver locations: %w", err)
	}
	defer cursor.Close(ctx)

	var locations []models.DriverLocation
	if err := cursor.All(ctx, &locations); err != nil {
		return nil, fmt.Errorf("failed to decode driver locations: %w", err)
	}

	return locations, nil
}

