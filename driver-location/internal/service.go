package internal

import "context"

type LocationDriverService struct {
	repo Repository
}

func NewLocationDriverService() *LocationDriverService {
	client, err := OpenConnection()
	if err != nil {
		panic(err)
	}
	
	repository := NewDriverLocationRepository(client.Database("bitaksi"), "locations")
	return &LocationDriverService{
		repo: repository,
	}
}

func (s *LocationDriverService) CreateDriverLocation(longitude, latitude float64) error {
	driver_location := DriverLocation{
		Location: Location{
			Type: "Point",
			Coordinates: []float64{longitude, latitude},
		},
	}
	return s.repo.CreateDriverLocation(context.TODO(), &driver_location)
}


func (s *LocationDriverService) BulkCreateDriverLocations(locations [][]float64) error {
	driver_locations := make([]*DriverLocation, len(locations))
	for i, location := range locations {
		driver_locations[i] = &DriverLocation{
			Location: Location{
				Type: "Point",
				Coordinates: location,
			},
		}
	}
	return s.repo.BulkCreateDriverLocations(context.TODO(), driver_locations)
}

func (s *LocationDriverService) SearchLocation(longitude, latitude float64, radius int) ([]DriverLocation, error) {
	return s.repo.SearchLocation(context.TODO(), longitude, latitude, radius)
}