package suite

import (
	"fmt"
	"gopoc/templates"
	endpoints "gopoc/utils/endpoints"
	helpers "gopoc/utils/helpers"
	"net/http"
	"testing"

	"github.com/bxcodec/faker/v4"
	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type UserTestSuite struct{
	suite.Suite
}
func (s *UserTestSuite) TestCreateUser (t provider.T) {

	t.Epic("User Management")
	t.Feature("User Creation")
	t.Story("Create a new user")

	var user map[string]interface{} 

	t.WithNewStep("Generate fake user data", func (sCtx provider.StepCtx)  {
		

		helpers.ConvertJsonTemplateToMap(templates.UserCreateTemplate, &user)

		user["name"] = faker.Name()
		user["last_name"] = faker.LastName()
		user["email"] = faker.Email()

		sCtx.WithNewAttachment("Generated User Data", allure.Text, []byte(fmt.Sprintf("%v", user)))
	})

	t.WithNewStep("Send Create User Request", func(sCtx provider.StepCtx) {
		response := endpoints.CreateUser(&testing.T{} ,user)

		sCtx.WithNewAttachment("Response", allure.Text, []byte(fmt.Sprintf("%v", response.Body())))

		sCtx.NewStep("Validate Response")
		sCtx.Assert().Equal("201 Created", response.Raw().Status)
		

		sCtx.NewStep("Validate Id is not null")
		sCtx.Assert().NotNil(response.JSON().Object().Value("id").Raw())
		
	})
	

}

func (s *UserTestSuite) TestGetUsers(t provider.T) {

	t.Epic("User Management")
	t.Feature("User Retrieval")
	t.Story("Get list of users")

	t.WithNewStep("Send Get Users Request", func(sCtx provider.StepCtx) {
		response := endpoints.GetUsers(&testing.T{})

		sCtx.WithNewAttachment("Response", allure.Text, []byte(fmt.Sprintf("%v", response.Body())))

		sCtx.NewStep("Validate Response")
		sCtx.Assert().Equal(http.StatusOK, response.Raw().StatusCode)

		sCtx.NewStep("Validate Response Schema")
		
	})
}

func (s *UserTestSuite) TestGetUserById(t provider.T) {

	t.Epic("User Management")
	t.Feature("User Retrieval")
	t.Story("Get user by ID")

	var user map[string]interface{}

	t.WithNewStep("Generate fake user data", func(sCtx provider.StepCtx) {
		helpers.ConvertJsonTemplateToMap(templates.UserCreateTemplate, &user)

		user["name"] = faker.Name()
		user["last_name"] = faker.LastName()
		user["email"] = faker.Email()

		sCtx.WithNewAttachment("Generated User Data", allure.Text, []byte(fmt.Sprintf("%v", user)))
	})

	var userId int

	t.WithNewStep("Create a new user", func(sCtx provider.StepCtx) {
		createResponse := endpoints.CreateUser(&testing.T{}, user)

		sCtx.WithNewAttachment("Create Response", allure.Text, []byte(fmt.Sprintf("%v", createResponse.Body())))

		sCtx.NewStep("Validate Create Response")
		sCtx.Assert().Equal(http.StatusCreated, createResponse.Raw().StatusCode)

		userId = int(createResponse.JSON().Object().Value("id").Number().Raw())
	})

	t.WithNewStep("Retrieve user by ID", func(sCtx provider.StepCtx) {
		getResponse := endpoints.GetUserById(&testing.T{}, userId)
		userObject := getResponse.JSON().Object()

		sCtx.WithNewAttachment("Get Response", allure.Text, []byte(fmt.Sprintf("%v", getResponse.Body())))

		sCtx.NewStep("Validate Get Response")
		sCtx.Assert().Equal(http.StatusOK, getResponse.Raw().StatusCode)

		sCtx.NewStep("Validate User Data")
		sCtx.Assert().Equal(user["name"], userObject.Value("name").Raw())
		sCtx.Assert().Equal(user["last_name"], userObject.Value("last_name").Raw())
		sCtx.Assert().Equal(user["email"], userObject.Value("email").Raw())
	})
}

func (s *UserTestSuite) TestUpdateUser(t provider.T) {

	t.Epic("User Management")
	t.Feature("User Update")
	t.Story("Update user details")

	var user map[string]interface{}
	var userUpdated map[string]interface{}

	t.WithNewStep("Generate fake user data", func(sCtx provider.StepCtx) {
		helpers.ConvertJsonTemplateToMap(templates.UserCreateTemplate, &user)

		user["name"] = faker.Name()
		user["last_name"] = faker.LastName()
		user["email"] = faker.Email()

		sCtx.WithNewAttachment("Generated User Data", allure.Text, []byte(fmt.Sprintf("%v", user)))
	})

	var userId int

	t.WithNewStep("Create a new user", func(sCtx provider.StepCtx) {
		createResponse := endpoints.CreateUser(&testing.T{}, user)

		sCtx.WithNewAttachment("Create Response", allure.Text, []byte(fmt.Sprintf("%v", createResponse.Body())))

		sCtx.NewStep("Validate Create Response")
		sCtx.Assert().Equal(http.StatusCreated, createResponse.Raw().StatusCode)

		userId = int(createResponse.JSON().Object().Value("id").Number().Raw())
	})

	t.WithNewStep("Generate updated user data", func(sCtx provider.StepCtx) {
		helpers.ConvertJsonTemplateToMap(templates.UserUpdateTemplate, &userUpdated)

		userUpdated["name"] = faker.Name()
		userUpdated["last_name"] = faker.LastName()
		userUpdated["email"] = faker.Email()

		sCtx.WithNewAttachment("Updated User Data", allure.Text, []byte(fmt.Sprintf("%v", userUpdated)))
	})

	t.WithNewStep("Update user by ID", func(sCtx provider.StepCtx) {
		updateResponse := endpoints.UpdateUserById(&testing.T{}, userId, userUpdated)

		sCtx.WithNewAttachment("Update Response", allure.Text, []byte(fmt.Sprintf("%v", updateResponse.Body())))

		sCtx.NewStep("Validate Update Response")
		sCtx.Assert().Equal(http.StatusOK, updateResponse.Raw().StatusCode)
	})

	t.WithNewStep("Retrieve updated user by ID", func(sCtx provider.StepCtx) {
		getResponse := endpoints.GetUserById(&testing.T{}, userId)
		userObject := getResponse.JSON().Object()

		sCtx.WithNewAttachment("Get Response", allure.Text, []byte(fmt.Sprintf("%v", getResponse.Body())))

		sCtx.NewStep("Validate Get Response")
		sCtx.Assert().Equal(http.StatusOK, getResponse.Raw().StatusCode)

		sCtx.NewStep("Validate Updated User Data")
		sCtx.Assert().Equal(userUpdated["name"], userObject.Value("name").Raw())
		sCtx.Assert().Equal(userUpdated["last_name"], userObject.Value("last_name").Raw())
		sCtx.Assert().Equal(userUpdated["email"], userObject.Value("email").Raw())
	})
}

func (s *UserTestSuite) TestDeleteUser(t provider.T) {

	t.Epic("User Management")
	t.Feature("User Deletion")
	t.Story("Delete a user by ID")

	var user map[string]interface{}

	t.WithNewStep("Generate fake user data", func(sCtx provider.StepCtx) {
		helpers.ConvertJsonTemplateToMap(templates.UserCreateTemplate, &user)

		user["name"] = faker.Name()
		user["last_name"] = faker.LastName()
		user["email"] = faker.Email()

		sCtx.WithNewAttachment("Generated User Data", allure.Text, []byte(fmt.Sprintf("%v", user)))
	})

	var userId int

	t.WithNewStep("Create a new user", func(sCtx provider.StepCtx) {
		createResponse := endpoints.CreateUser(&testing.T{}, user)

		sCtx.WithNewAttachment("Create Response", allure.Text, []byte(fmt.Sprintf("%v", createResponse.Body())))

		sCtx.NewStep("Validate Create Response")
		sCtx.Assert().Equal(http.StatusCreated, createResponse.Raw().StatusCode)

		userId = int(createResponse.JSON().Object().Value("id").Number().Raw())
	})

	t.WithNewStep("Delete user by ID", func(sCtx provider.StepCtx) {
		deleteResponse := endpoints.DeleteUserById(&testing.T{}, userId)

		sCtx.WithNewAttachment("Delete Response", allure.Text, []byte(fmt.Sprintf("%v", deleteResponse.Body())))

		sCtx.NewStep("Validate Delete Response")
		sCtx.Assert().Equal(http.StatusOK, deleteResponse.Raw().StatusCode)
	})

	t.WithNewStep("Attempt to retrieve deleted user by ID", func(sCtx provider.StepCtx) {
		getResponse := endpoints.GetUserById(&testing.T{}, userId)

		sCtx.WithNewAttachment("Get Response", allure.Text, []byte(fmt.Sprintf("%v", getResponse.Body())))

		sCtx.NewStep("Validate Get Response for Deleted User")
		sCtx.Assert().Equal(http.StatusNotFound, getResponse.Raw().StatusCode)
		sCtx.Assert().Equal("User not found", getResponse.JSON().Object().Value("message").Raw())
	})
}

func TestUser(t *testing.T){
	t.Parallel()

	suite.RunSuite(t, new(UserTestSuite))
}