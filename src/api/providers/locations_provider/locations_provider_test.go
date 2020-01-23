package locations_provider_test

import (
	"fmt"
	"testing"

	. "golang_testing/src/api/providers/locations_provider"
	"net/http"

	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
)

func TestGetCountryRestClientError(t *testing.T) {
	rest.StartMockupServer()
	rest.AddMockups(&rest.Mock{
		URL:          "https://api.mercadolibre.com/countries/AR",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: 0,
	})

	country, err := GetCountry("AR")
	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid restclient error when getting country AR", err.Message)

}

func TestGetCountryNotFound(t *testing.T) {
	rest.StartMockupServer()

	rest.AddMockups(&rest.Mock{
		URL:          "https://api.mercadolibre.com/countries/AR",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message": "Country not found", "error", "not_found", "status" : 404, "cause":[]}`,
	})
	countryID := "ARS"
	country, err := GetCountry(countryID)
	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	errMessage := fmt.Sprintf("Country not found")
	assert.EqualValues(t, errMessage, err.Message)
}

func TestGetCountryInvalidErrorInterface(t *testing.T) {
	country, err := GetCountry("AR")
	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid error interface when getting country AR", err.Message)
}

func TestGetCountryInvalidJsonResponse(t *testing.T) {
	country, err := GetCountry("AR")
	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "error when trying to unmarshal country data for AR", err.Message)
}

func TestGetCountryNoError(t *testing.T) {
	country, err := GetCountry("AR")
	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, "AR", country.Id)
	assert.EqualValues(t, "Argentina", country.Name)
	assert.EqualValues(t, "GMT-03:00", country.TimeZone)
	assert.EqualValues(t, 24, len(country.States))
}
