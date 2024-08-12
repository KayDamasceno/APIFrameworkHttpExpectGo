package utils

import (
	"testing"

	"github.com/gavv/httpexpect/v2"
)



func SendQueryCountries(t *testing.T, query string) *httpexpect.Response {

	e := httpexpect.Default(t, countriesURL)

	response := e.POST("/").
				WithHeader("Content-Type", "application/json").
				WithJSON(map[string] interface{} {
					"query": query,
				})
	
	return response.Expect()
				

}