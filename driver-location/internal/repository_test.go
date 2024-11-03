package internal_test

import (
	"context"
	"testing"

	"github.com/cinarizasyon/bitaksi-golang-case-study/driver-location/internal"
	"github.com/cinarizasyon/bitaksi-golang-case-study/driver-location/internal/models"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestCreateDriverLocation(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("CreateDriverLocation_WhenResponseIsSuccess", func(mt *mtest.T) {
		driverLocationRepo := internal.NewDriverLocationRepository(mt.DB, "locations")
		mt.AddMockResponses(mtest.CreateSuccessResponse())

		err := driverLocationRepo.CreateDriverLocation(context.TODO(), &models.DriverLocation{
			Location: models.Location{
				Type:        "Point",
				Coordinates: []float64{41.0082, 28.9784},
			},
		})

		assert.Nil(t, err)
	})

	mt.Run("CreateDriverLocation_WhenResponseIsFail", func(mt *mtest.T) {
		driverLocationRepo := internal.NewDriverLocationRepository(mt.DB, "locations")
		mt.AddMockResponses(bson.D{{Key: "ok", Value: -1}})

		err := driverLocationRepo.CreateDriverLocation(context.TODO(), &models.DriverLocation{
			Location: models.Location{
				Type:        "Point",
				Coordinates: []float64{41.0082, 28.9784},
			},
		})

		assert.NotNil(t, err)
	})
}

func TestBulkCreateDriverLocations(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("BulkCreateDriverLocations_WhenResponseIsSuccess", func(mt *mtest.T) {
		driverLocationRepo := internal.NewDriverLocationRepository(mt.DB, "locations")
		mt.AddMockResponses(mtest.CreateSuccessResponse())

		err := driverLocationRepo.BulkCreateDriverLocations(context.TODO(), []*models.DriverLocation{
			{
				Location: models.Location{
					Type:        "Point",
					Coordinates: []float64{41.0082, 28.9784},
				},
			},
			{
				Location: models.Location{
					Type:        "Point",
					Coordinates: []float64{42.0082, 21.9784},
				},
			},
		})

		assert.Nil(t, err)
	})

	mt.Run("BulkCreateDriverLocations_WhenResponseIsFail", func(mt *mtest.T) {
		driverLocationRepo := internal.NewDriverLocationRepository(mt.DB, "locations")
		mt.AddMockResponses(bson.D{{Key: "ok", Value: -1}})

		err := driverLocationRepo.BulkCreateDriverLocations(context.TODO(), []*models.DriverLocation{
			{
				Location: models.Location{
					Type:        "Point",
					Coordinates: []float64{41.0082, 28.9784},
				},
			},
			{
				Location: models.Location{
					Type:        "Point",
					Coordinates: []float64{42.0082, 21.9784},
				},
			},
		})

		assert.NotNil(t, err)
	})
}

func TestSearchLocation(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("SearchLocation_WhenResponseIsSuccess", func(mt *mtest.T) {
		ns := mt.Coll.Database().Name() + "." + mt.Coll.Name()
		driverLocationRepo := internal.NewDriverLocationRepository(mt.DB, mt.Coll.Name())
		mt.AddMockResponses(mtest.CreateCursorResponse(1, ns, mtest.FirstBatch,
			bson.D{
				{Key: "_id", Value: "5f5c5b0d3b4b3f1c9f4b8e8d"},
				{Key: "location", Value: bson.D{
					{Key: "type", Value: "Point"},
					{Key: "coordinates", Value: bson.A{41.0082, 28.9784}},
				}},
				{Key: "distance", Value: 0.0},
			},
			bson.D{
				{Key: "_id", Value: "5d5c5b0d3baf3f1c9f4b8e8f"},
				{Key: "location", Value: bson.D{
					{Key: "type", Value: "Point"},
					{Key: "coordinates", Value: bson.A{41.182, 28.9784}},
				}},
				{Key: "distance", Value: 0.0},
			},
		))

		searchResult, err := driverLocationRepo.SearchLocation(context.TODO(), 41.0082, 28.9784, 1000)

		assert.Nil(t, err)
		assert.Equal(t, 2, len(searchResult))

		assert.Equal(t, "5f5c5b0d3b4b3f1c9f4b8e8d", searchResult[0].Lookup("_id").StringValue())
		assert.Equal(t, 0.0, searchResult[0].Lookup("distance").Double())
		assert.Equal(t, 41.0082, searchResult[0].Lookup("location").Document().Lookup("coordinates").Array().Index(0).Value().Double())
		assert.Equal(t, 28.9784, searchResult[0].Lookup("location").Document().Lookup("coordinates").Array().Index(1).Value().Double())
	
		assert.Equal(t, "5d5c5b0d3baf3f1c9f4b8e8f", searchResult[1].Lookup("_id").StringValue())
		assert.Equal(t, 0.0, searchResult[1].Lookup("distance").Double())
		assert.Equal(t, 41.182, searchResult[1].Lookup("location").Document().Lookup("coordinates").Array().Index(0).Value().Double())
		assert.Equal(t, 28.9784, searchResult[1].Lookup("location").Document().Lookup("coordinates").Array().Index(1).Value().Double())
	})

	mt.Run("SearchLocation_WhenResponseIsEmpty", func(mt *mtest.T) {
		ns := mt.Coll.Database().Name() + "." + mt.Coll.Name()
		driverLocationRepo := internal.NewDriverLocationRepository(mt.DB, mt.Coll.Name())
		mt.AddMockResponses(mtest.CreateCursorResponse(0, ns, mtest.FirstBatch))

		searchResult, err := driverLocationRepo.SearchLocation(context.TODO(), 41.0082, 28.9784, 1000)

		assert.Nil(t, err)
		assert.Equal(t, 0, len(searchResult))
	})

	mt.Run("SearchLocation_WhenResponseIsFail", func(mt *mtest.T) {
		driverLocationRepo := internal.NewDriverLocationRepository(mt.DB, "locations")
		mt.AddMockResponses(bson.D{{Key: "ok", Value: -1}})

		searchResult, err := driverLocationRepo.SearchLocation(context.TODO(), 41.0082, 28.9784, 1000)

		assert.NotNil(t, err)
		assert.Nil(t, searchResult)
	})
}
