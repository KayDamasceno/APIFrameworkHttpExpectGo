package utils

import "encoding/json"

func ConvertJsonTemplateToMap(template string, v interface{}) error {
	return json.Unmarshal([]byte(template), v)
}