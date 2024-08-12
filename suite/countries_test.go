package suite

import (
	"gopoc/templates"
	countries "gopoc/utils/endpoints"
	"net/http"
	"testing"
)

func TestGetBrazil(t *testing.T) {

	queryResponse := countries.SendQueryCountries(t, templates.GetBrazilQuery)

	queryResponse.Status(http.StatusOK)

	
	brazilObject := queryResponse.JSON().Object().Value("data").Object().Value("country").Object()
	
	brazilObject.Value("name").IsEqual("Brazil")
	brazilObject.Value("native").IsEqual("Brasil")
	brazilObject.Value("capital").IsEqual("Brasília")
	brazilObject.Value("emoji").IsEqual("🇧🇷")
	brazilObject.Value("currency").IsEqual("BRL")

	
	languagesArray := brazilObject.Value("languages").Array()
	
	languagesArray.Value(0).Object().Value("code").IsEqual("pt")
	languagesArray.Value(0).Object().Value("name").IsEqual("Portuguese")
	
}