package suite

import (
	"gopoc/schemas"
	"gopoc/templates"
	endpoints "gopoc/utils/endpoints"
	helpers "gopoc/utils/helpers"
	"testing"

	"github.com/bxcodec/faker/v4"
	"github.com/smartystreets/goconvey/convey"
)

func TestCreateUser (t *testing.T) {

	convey.Convey("Given a valid user creation request", t, func ()  {
			
		var user map[string]interface{} 

		helpers.ConvertJsonTemplateToMap(templates.UserCreateTemplate, &user)

		user["name"] = faker.Name()
		user["last_name"] = faker.LastName()
		user["email"] = faker.Email()

		convey.Convey("When the POST request is made passing the user payload", func() {
			response := endpoints.CreateUser(t, user)

			convey.Convey("Then the response status should be 201 Created", func ()  {
				convey.So(response.Raw().StatusCode, convey.ShouldEqual, 201)

				convey.Convey("And the response should contain a non null id", func() {
					convey.So(response.JSON().Object().Value("id").Raw(), convey.ShouldNotBeNil)
				})
			})
		})	
	})

}
func TestGetUsers(t *testing.T){

	convey.Convey("Given the system have users", t, func() {
		
		convey.Convey("When the GET request is made to return all the users", func() {
			response := endpoints.GetUsers(t)
			
			convey.Convey("Then the response status should be 200 Ok", func() {
				convey.So(response.Raw().StatusCode, convey.ShouldEqual, 200)

				convey.Convey("And the response should match the schema", func() {
					response.JSON().Array().Schema(schemas.UsersGetSchema)
				})
			})
		})
	})

	
	
}

func TestGetUserById(t *testing.T){

	convey.Convey("Given a you have a valid user with an identifier", t, func ()  {
		var user map[string]interface{} 

		helpers.ConvertJsonTemplateToMap(templates.UserCreateTemplate, &user)

		user["name"] = faker.Name()
		user["last_name"] = faker.LastName()
		user["email"] = faker.Email()

		createResponse := endpoints.CreateUser(t, user)
		id := createResponse.JSON().Object().Value("id").Number().Raw()

		convey.So(createResponse.Raw().StatusCode, convey.ShouldEqual, 201)
		
		convey.Convey("When a GET request is made passing the identifier number", func ()  {
			getResponse := endpoints.GetUserById(t, int(id))
			userObject := getResponse.JSON().Object()

			convey.Convey("Then the response status should be 200 Ok", func ()  {
				convey.So(getResponse.Raw().StatusCode, convey.ShouldEqual, 200)

				convey.Convey("And the body should return the info from the user requested", func() {
					convey.So(userObject.Value("name").Raw(), convey.ShouldEqual, user["name"])
					convey.So(userObject.Value("last_name").Raw(), convey.ShouldEqual, user["last_name"])
					convey.So(userObject.Value("email").Raw(), convey.ShouldEqual, user["email"])
				})
			})
		})
	})

	

}

func TestUpdateUser(t *testing.T) {

	convey.Convey("Given you have a user with an identifier", t, func ()  {
		
		var user map[string]interface{}
		var userUpdated map[string]interface{}

		helpers.ConvertJsonTemplateToMap(templates.UserCreateTemplate, &user)
		helpers.ConvertJsonTemplateToMap(templates.UserUpdateTemplate, &userUpdated)

		user["name"] = faker.Name()
		user["last_name"] = faker.LastName()
		user["email"] = faker.Email()

		createResponse := endpoints.CreateUser(t, user)
		id := createResponse.JSON().Object().Value("id").Number().Raw()

		convey.So(createResponse.Raw().StatusCode, convey.ShouldEqual, 201)

		convey.Convey("When a PUT request is made to this identifier passing updates about the user", func ()  {
			
			userUpdated["name"] = faker.Name()
			userUpdated["last_name"] = faker.LastName()
			userUpdated["email"] = faker.Email()

			updateResponse := endpoints.UpdateUserById(t, int(id), userUpdated)

			convey.Convey("Then the response status should be 200 Ok", func ()  {
				convey.So(updateResponse.Raw().StatusCode, convey.ShouldEqual, 200)

				convey.Convey("And the user should be updated with the new values", func() {
					getResponse := endpoints.GetUserById(t, int(id))
					userObject := getResponse.JSON().Object()

					convey.So(getResponse.Raw().StatusCode, convey.ShouldEqual, 200)
					convey.So(userObject.Value("name").Raw(), convey.ShouldEqual, userUpdated["name"])
					convey.So(userObject.Value("last_name").Raw(), convey.ShouldEqual, userUpdated["last_name"])
					convey.So(userObject.Value("email").Raw(), convey.ShouldEqual, userUpdated["email"])

				})
			})
		})

	})

	
	

}

func TestDeleteUser(t *testing.T){

	convey.Convey("Given you have an user with an identifer", t, func ()  {
		var user map[string]interface{} 

		helpers.ConvertJsonTemplateToMap(templates.UserCreateTemplate, &user)

		user["name"] = faker.Name()
		user["last_name"] = faker.LastName()
		user["email"] = faker.Email()


		createResponse := endpoints.CreateUser(t, user)
		id := createResponse.JSON().Object().Value("id").Number().Raw()


		convey.So(createResponse.Raw().StatusCode, convey.ShouldEqual, 201)

		convey.Convey("When a DELETE request is made to the delete the user passing the identifier", func ()  {
			deleteResponse := endpoints.DeleteUserById(t, int(id))

			convey.Convey("Then the response status should be 200 Ok", func ()  {
				convey.So(deleteResponse.Raw().StatusCode, convey.ShouldEqual, 200)

				convey.Convey("And the user should not be more available in the system", func() {
					getResponse := endpoints.GetUserById(t, int(id))
					convey.So(getResponse.Raw().StatusCode, convey.ShouldEqual, 404)

				})
			})

		})

	})


}