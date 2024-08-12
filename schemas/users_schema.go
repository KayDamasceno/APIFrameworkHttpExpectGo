package schemas

const UsersGetSchema = `{
	"type":"array",
	"items": {
		"type":"object",
		"properties":{
			"id": {"type": "integer"},
			"email": {"type": "string"},
			"last_name": {"type": "string"},
			"name": {"type": "string"}
		},
		"required": ["id", "email", "last_name", "name"]
	}
}`