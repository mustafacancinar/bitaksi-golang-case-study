package internal

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

func InitDatabase() error {
	client, err := OpenConnection()
	if err != nil {
		return fmt.Errorf("failed to open connection: %w", err)
	}
	defer CloseConnection(client)

	database := client.Database("bitaksi")

	validation := bson.M{
		"$jsonSchema": bson.M{
			"bsonType": "object",
			"required": []string{"location"},
			"properties": bson.M{
				"location": bson.M{
					"bsonType": "object",
					"required": []string{"type", "coordinates"},
					"properties": bson.M{
						"type": bson.M{
							"bsonType": "string",
							"enum": []string{"Point"},
						},
						"coordinates": bson.M{
							"bsonType": "array",
							"minItems": 2,
							"maxItems": 2,
							"items": []bson.M{
								{"bsonType": "double", "minimum": -180, "maximum": 180},
								{"bsonType": "double", "minimum": -90, "maximum": 90},
							},
						},
					},
				},
			},
		},
	}
	opts := options.CreateCollection().SetValidator(validation)

	err = database.CreateCollection(context.Background(), "locations", opts)
	if err != nil {
		return fmt.Errorf("failed to create collection: %w", err)
	}

	collection  := database.Collection("locations")
	_, err = collection.Indexes().CreateOne(
		context.Background(), 
		mongo.IndexModel{
			Keys: bson.D{{Key: "location", Value: "2dsphere"}},
		},
	)
	if err != nil {
		return fmt.Errorf("failed to create index: %w", err)
	}

	return nil
}

func OpenConnection() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	mongoClient, err := mongo.Connect(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %w", err)
	}
	
	if err := mongoClient.Ping(ctx, readpref.Primary()); err != nil {
		return nil, fmt.Errorf("failed to ping server: %w", err)
	}

	return mongoClient, nil
}

func CloseConnection(client *mongo.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Disconnect(ctx); err != nil {
		return fmt.Errorf("failed to close connection: %w", err)
	}

	return nil
}

func Ping(client *mongo.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return fmt.Errorf("failed to ping server: %w", err)
	}

	return nil
}