package internal

import (
	"bytes"
	"encoding/json"
	"net"
	"net/http"
	"time"
)

type MatchingService struct {
	DriverLocationServiceBaseUrl string
}

func NewMatchingService(driverLocationServiceBaseUrl string) *MatchingService {
	return &MatchingService{
		DriverLocationServiceBaseUrl: driverLocationServiceBaseUrl,
	}
}

func (s *MatchingService) Match(request *MatchingRequest) (*MatchingResponse, error) {
	payload, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	url := s.DriverLocationServiceBaseUrl + "/drivers/search"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", request.Token)

	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, nil
	}

	var matchingResponse []MatchingResponse
	err = json.NewDecoder(response.Body).Decode(&matchingResponse)
	if err != nil {
		return nil, err
	}

	if len(matchingResponse) == 0 {
		return nil, nil
	}

	return &matchingResponse[0], nil
}

func (s *MatchingService) CheckRemoteServiceHealth() (bool, error) {
	url := s.DriverLocationServiceBaseUrl + "/healthz"
	client:= &http.Client{
		Timeout: 5 * time.Second,
	}
	
	response, err := client.Get(url)
	if err != nil {
		netErr, ok := err.(net.Error)
		if ok && netErr.Timeout() {
			return false, nil
		}
		return false, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return false, nil
	}

	return true, nil
}







