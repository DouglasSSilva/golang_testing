package services

import (
	"golang_testing/src/api/domain/locations"
	"golang_testing/src/api/providers/locations_provider"
	"golang_testing/src/api/utils/errors"
)

func GetCountry(countryID string) (*locations.Country, *errors.APIError) {
	return locations_provider.GetCountry(countryID)
}
