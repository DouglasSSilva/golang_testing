package locations_provider

import (
	"encoding/json"
	"fmt"
	"net/http"

	"golang_testing/src/api/domain/locations"
	"golang_testing/src/api/utils/errors"

	"github.com/mercadolibre/golang-restclient/rest"
)

const (
	urlGetCountry = "https://api.mercadolibre.com/countries/%s"
)

func GetCountry(countryID string) (*locations.Country, *errors.APIError){
	response := rest.Get(fmt.Sprintf(urlGetCountry, countryID))
	if response == nil || response.Response == nil {
		return nil, &errors.APIError{	
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("invalid restclient error when getting country %s", countryID),
		}
	}
	fmt.Println("Status Code", response.StatusCode)
	if response.StatusCode > 299 {
		var apiErr errors.APIError
		if err := json.Unmarshal(response.Bytes(), &apiErr); err == nil {
			return nil, &errors.APIError{
				Status:  http.StatusInternalServerError,
				Message: fmt.Sprintf("invalid error response when getting country %s", countryID),
			}
		}
	}

	var result locations.Country

	if err := json.Unmarshal(response.Bytes(), &result); err != nil {
		return nil, &errors.APIError{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("error when  trying to unmarshal country data %s", countryID),
		}
	}

	return &result, nil
}
