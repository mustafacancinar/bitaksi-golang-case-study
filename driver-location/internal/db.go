package internal

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)


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