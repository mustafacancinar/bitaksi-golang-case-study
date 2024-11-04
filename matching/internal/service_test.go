package internal_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cinarizasyon/bitaksi-golang-case-study/matching/internal"
	"github.com/stretchr/testify/assert"
)

func TestMatch(t *testing.T) {
	assert := assert.New(t)

	t.Run("When remote service returns a valid response", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(`[{"id":"6727bcf4fb1a48d03b8b7b87","latitude":41.0082,"longitude":28.9784, "distance": 0.7982103583486942}]`))
		}))
	
		defer server.Close()
	
		service := internal.NewMatchingService(server.URL)
	
		request := &internal.MatchingRequest{
			Longitude: 28.9784,
			Latitude:  41.0082,
			Radius:    1,
		}
	
		response, err := service.Match(request)
		assert.Nil(err)
		assert.Equal("6727bcf4fb1a48d03b8b7b87", response.DriverId)
		assert.Equal(0.7982103583486942, response.Distance)
		assert.Equal(28.9784, response.Longitude)
		assert.Equal(41.0082, response.Latitude)
	})

	t.Run("When remote service returns an empty response", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(`[]`))
		}))
	
		defer server.Close()
	
		service := internal.NewMatchingService(server.URL)
	
		request := &internal.MatchingRequest{
			Longitude: 28.9784,
			Latitude:  41.0082,
			Radius:    1,
		}
	
		response, err := service.Match(request)
		assert.Nil(err)
		assert.Nil(response)
	})

	t.Run("When remote service returns an invalid response", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(`[{abc}]`))
		}))
	
		defer server.Close()
	
		service := internal.NewMatchingService(server.URL)
	
		request := &internal.MatchingRequest{
			Longitude: 28.9784,
			Latitude:  41.0082,
			Radius:    1,
		}
	
		response, err := service.Match(request)
		assert.NotNil(err)
		assert.Nil(response)
		assert.Equal("invalid character 'a' looking for beginning of object key string", err.Error())
	})
}

func TestCheckRemoteServiceHealth(t *testing.T) {
	assert := assert.New(t)

	t.Run("When remote service is healthy", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))
	
		defer server.Close()
	
		service := internal.NewMatchingService(server.URL)
	
		isHealthy, err := service.CheckRemoteServiceHealth()
		assert.Nil(err)
		assert.True(isHealthy)
	})

	t.Run("When remote service is not healthy", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusServiceUnavailable)
		}))
	
		defer server.Close()
	
		service := internal.NewMatchingService(server.URL)
	
		isHealthy, err := service.CheckRemoteServiceHealth()
		assert.Nil(err)
		assert.False(isHealthy)
	})
}