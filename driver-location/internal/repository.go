package internal

import (
	"context"
	"fmt"

	"github.com/cinarizasyon/bitaksi-golang-case-study/driver-location/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	CreateDriverLocation(ctx context.Context, location *models.DriverLocation) error
	BulkCreateDriverLocations(ctx context.Context, locations []*models.DriverLocation) error
	SearchLocation(ctx context.Context, longitude, latitude float64, radius int) ([]bson.Raw, error)
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
	models := make([]mongo.WriteModel, len(locations))
	for i, location := range locations {
		models[i] = mongo.NewInsertOneModel().SetDocument(location)
	}

	_, err := r.collection.BulkWrite(ctx, models)
	if err != nil {
		return fmt.Errorf("failed to insert driver locations: %w", err)
	}

	return nil
}

func (r *DriverLocationRepository) SearchLocation(ctx context.Context, longitude, latitude float64, radius int) ([]bson.Raw, error) {
	pipeline := mongo.Pipeline{
		{
			{Key: "$geoNear", Value: bson.D{
				{Key: "near", Value: bson.D{
					{Key: "type", Value: "Point"},
					{Key: "coordinates", Value: []float64{longitude, latitude}},
				}},
				{Key: "distanceField", Value: "distance"},
				{Key: "spherical", Value: true},
				{Key: "maxDistance", Value: float64(radius) * 1000},
				{Key: "distanceMultiplier", Value: 0.001},
			}},
		},
		{
			{Key: "$limit", Value: 1},
		},
	}

	cursor, err := r.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("failed to find driver locations: %w", err)
	}
	defer cursor.Close(ctx)

	var locations []bson.Raw
	for cursor.Next(ctx) {
		var location bson.Raw

		if err := cursor.Decode(&location); err != nil {
			return nil, fmt.Errorf("failed to decode driver location: %w", err)
		}

		locations = append(locations, location)
	}

	return locations, nil
}
