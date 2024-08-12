package utils

import (
	"testing"

	"github.com/gavv/httpexpect/v2"
)


func CreateUser(t *testing.T, payload interface{} ) *httpexpect.Response {
	
	api := httpexpect.Default(t, desafioQaURL)

	postResponse := api.POST("/users").WithJSON(payload)

	return postResponse.Expect()
	
}


func GetUsers (t *testing.T) *httpexpect.Response {

	api := httpexpect.Default(t, desafioQaURL)

	getResponse := api.GET("/users")
	
	return getResponse.Expect()
}


func GetUserById (t *testing.T, id int) *httpexpect.Response {
	
	api := httpexpect.Default(t, desafioQaURL)

	getResponse := api.GET("/users/{id}").WithPath("id", id)

	return getResponse.Expect()
	
}

func UpdateUserById (t *testing.T, id int, updatedPayload interface{}) *httpexpect.Response {

	api := httpexpect.Default(t, desafioQaURL)

	putResponse := api.PUT("/users/{id}").WithPath("id", id).WithJSON(updatedPayload)

	return putResponse.Expect()
}

func DeleteUserById (t *testing.T, id int) *httpexpect.Response {
	
	api := httpexpect.Default(t, desafioQaURL)

	deleteResponse := api.DELETE("/users/{id}").WithPath("id", id)

	return deleteResponse.Expect()
	
}