package suite

import (
	"gopoc/templates"
	countries "gopoc/utils/endpoints"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestGetBrazil(t *testing.T) {

	convey.Convey("Given you have a query request to return a specific country", t, func ()  {
		convey.Convey("When the POST request is made for Brazil country", func ()  {
			queryResponse := countries.SendQueryCountries(t, templates.GetBrazilQuery)

			convey.Convey("Then the response status should be equal 200 Ok", func ()  {
				convey.So(queryResponse.Raw().StatusCode, convey.ShouldEqual, 200)

				convey.Convey("And the response should contain the correct information from Brazil", func ()  {
					
					brazilObject := queryResponse.JSON().Object().Value("data").Object().Value("country").Object()

					convey.So(brazilObject.Value("name").Raw(), convey.ShouldEqual, "Brazil")
					convey.So(brazilObject.Value("capital").Raw(), convey.ShouldEqual, "Brasília")
					convey.So(brazilObject.Value("emoji").Raw(), convey.ShouldEqual, "🇧🇷")
					convey.So(brazilObject.Value("currency").Raw(), convey.ShouldEqual, "BRL")
					convey.So(brazilObject.Value("native").Raw(), convey.ShouldEqual, "Brasil")


					languagesArray := brazilObject.Value("languages").Array()
					convey.So(languagesArray.Value(0).Object().Value("code").Raw(), convey.ShouldEqual, "pt")
					convey.So(languagesArray.Value(0).Object().Value("name").Raw(), convey.ShouldEqual, "Portuguese")
				})
			})
		})
	})

	

	
}