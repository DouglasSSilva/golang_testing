package services

import (
	"golang_testing/src/api/domain/locations"
	"golang_testing/src/api/providers/locations_provider"
	"golang_testing/src/api/utils/errors"
)

type locationService struct{}

type locationServiceInterface interface {
	GetCountry(countryID string) (*locations.Country, *errors.APIError)
}

var (
	LocationsService locationServiceInterface
)

func init() {
	LocationsService = &locationService{}
}

func (s *locationService) GetCountry(countryID string) (*locations.Country, *errors.APIError) {
	return locations_provider.GetCountry(countryID)
}
