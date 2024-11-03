package internal

import (
	"context"

	"github.com/cinarizasyon/bitaksi-golang-case-study/driver-location/internal/models"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type LocationDriverService struct {
	repo Repository
	Client *mongo.Client
}

func NewLocationDriverService() *LocationDriverService {
	client, err := OpenConnection()
	if err != nil {
		panic(err)
	}
	
	repository := NewDriverLocationRepository(client.Database("bitaksi"), "locations")
	return &LocationDriverService{
		repo: repository,
		Client: client,
	}
}

func (s *LocationDriverService) CreateDriverLocation(req *models.CreateDriverLocationRequest) error {
	driver_location := models.DriverLocation{
		Location: models.Location{
			Type: "Point",
			Coordinates: []float64{req.Longitude, req.Latitude},
		},
	}
	return s.repo.CreateDriverLocation(context.TODO(), &driver_location)
}


func (s *LocationDriverService) BulkCreateDriverLocations(req *models.BulkCreateDriverLocationRequest) error {
	driver_locations := make([]*models.DriverLocation, len(req.Locations))
	for i, location := range req.Locations {
		driver_locations[i] = &models.DriverLocation{
			Location: models.Location{
				Type: "Point",
				Coordinates: []float64{location.Longitude, location.Latitude},
			},
		}
	}
	return s.repo.BulkCreateDriverLocations(context.TODO(), driver_locations)
}

func (s *LocationDriverService) SearchLocation(longitude, latitude float64, radius int) ([]models.DriverLocation, error) {
	return s.repo.SearchLocation(context.TODO(), longitude, latitude, radius)
}