package controllers

import (
	"encoding/json"
	"golang_testing/src/api/domain/locations"
	"golang_testing/src/api/services"
	"golang_testing/src/api/utils/errors"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var (
	getCountryFunc func(coutnryID string) (*locations.Country, *errors.APIError)
)

func TestMain(m *testing.M) {
	// rest.StartMockupServer()
	os.Exit(m.Run())
}

type locationsServiceMock struct{}

func (*locationsServiceMock) GetCountry(countryID string) (*locations.Country, *errors.APIError) {
	return getCountryFunc(countryID)
}

func TestGetCountryNotFound(t *testing.T) {
	// Mock locationservice methods
	getCountryFunc = func(countryID string) (*locations.Country, *errors.APIError) {
		return nil, &errors.APIError{Status: http.StatusNotFound, Message: "Country not found"}
	}

	services.LocationsService = &locationsServiceMock{}

	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Request, _ = http.NewRequest(http.MethodGet, "", nil)
	GetCountry(c)
	assert.EqualValues(t, http.StatusNotFound, response.Code)

	var apiErr errors.APIError
	err := json.Unmarshal(response.Body.Bytes(), &apiErr)
	assert.Nil(t, err)

	assert.EqualValues(t, http.StatusNotFound, apiErr.Status)
	assert.EqualValues(t, "", apiErr.Message)

}
