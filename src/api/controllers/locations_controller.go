package controllers

import (
	"github.com/gin-gonic/gin"
	"golang_testing/src/api/services"
	"net/http"
)

func GetCountry(c *gin.Context) {
	country, err := services.LocationsService.GetCountry("country_id")
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, country)
}
