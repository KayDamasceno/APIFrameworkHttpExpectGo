package suite

import (
	"gopoc/schemas"
	"gopoc/templates"
	endpoints "gopoc/utils/endpoints"
	helpers "gopoc/utils/helpers"
	"net/http"
	"testing"

	"github.com/bxcodec/faker/v4"
)

func TestCreateUser (t *testing.T) {

	var user map[string]interface{} 

	helpers.ConvertJsonTemplateToMap(templates.UserCreateTemplate, &user)

	user["name"] = faker.Name()
	user["last_name"] = faker.LastName()
	user["email"] = faker.Email()

	response := endpoints.CreateUser(t, user)

	response.Status(http.StatusCreated)
	response.JSON().Object().Value("id").NotNull()

}
func TestGetUsers(t *testing.T){

	response := endpoints.GetUsers(t)

	response.Status(http.StatusOK)
	response.JSON().Array().Schema(schemas.UsersGetSchema)
}

func TestGetUserById(t *testing.T){

	var user map[string]interface{} 

	helpers.ConvertJsonTemplateToMap(templates.UserCreateTemplate, &user)

	user["name"] = faker.Name()
	user["last_name"] = faker.LastName()
	user["email"] = faker.Email()


	createResponse := endpoints.CreateUser(t, user)

	createResponse.Status(http.StatusCreated)

	id := createResponse.JSON().Object().Value("id").Number().Raw()

	getResponse := endpoints.GetUserById(t, int(id))
	userObject := getResponse.JSON().Object()

	getResponse.Status(http.StatusOK)
	userObject.Value("name").IsEqual(user["name"])
	userObject.Value("last_name").IsEqual(user["last_name"])
	userObject.Value("email").IsEqual(user["email"])
	

}

func TestUpdateUser(t *testing.T) {

	var user map[string]interface{}
	var userUpdated map[string]interface{}

	helpers.ConvertJsonTemplateToMap(templates.UserCreateTemplate, &user)
	helpers.ConvertJsonTemplateToMap(templates.UserUpdateTemplate, &userUpdated)

	user["name"] = faker.Name()
	user["last_name"] = faker.LastName()
	user["email"] = faker.Email()

	createResponse := endpoints.CreateUser(t, user)

	createResponse.Status(http.StatusCreated)

	id := createResponse.JSON().Object().Value("id").Number().Raw()

	userUpdated["name"] = faker.Name()
	userUpdated["last_name"] = faker.LastName()
	userUpdated["email"] = faker.Email()

	updateResponse := endpoints.UpdateUserById(t, int(id), userUpdated)
	
	updateResponse.Status(http.StatusOK)

	getResponse := endpoints.GetUserById(t, int(id))
	userObject := getResponse.JSON().Object()

	getResponse.Status(http.StatusOK)
	userObject.Value("name").IsEqual(userUpdated["name"])
	userObject.Value("last_name").IsEqual(userUpdated["last_name"])
	userObject.Value("email").IsEqual(userUpdated["email"])

}

func TestDeleteUser(t *testing.T){

	var user map[string]interface{} 

	helpers.ConvertJsonTemplateToMap(templates.UserCreateTemplate, &user)

	user["name"] = faker.Name()
	user["last_name"] = faker.LastName()
	user["email"] = faker.Email()


	createResponse := endpoints.CreateUser(t, user)

	createResponse.Status(http.StatusCreated)

	id := createResponse.JSON().Object().Value("id").Number().Raw()

	deleteResponse := endpoints.DeleteUserById(t, int(id))

	deleteResponse.Status(http.StatusOK)

	getResponse := endpoints.GetUserById(t, int(id))

	getResponse.Status(http.StatusNotFound)
	getResponse.JSON().Object().Value("message").IsEqual("User not found")

}