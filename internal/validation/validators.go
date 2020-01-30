package validation

import (
	"fmt"

	"github.com/xeipuuv/gojsonschema"
)

// ValidateURL validates URL.
func ValidateURL(URL string) (bool, error) {
	url := struct {
		URL string `json:"url"`
	}{
		URL: URL,
	}
	schema := gojsonschema.NewStringLoader(`
	{
		"$schema": "http://json-schema.org/draft-07/schema#",
		"type": "object",
		"properties": {
			"url": { 
				"type": "string", 
				"format": "uri",
    			"pattern": "^(http?|https?|)://"
			}
		}
	}`)
	loader := gojsonschema.NewGoLoader(url)

	return validate(schema, loader)
}

// ValidateUUID validates uuid.
func ValidateUUID(UUID string) (bool, error) {
	uuid := struct {
		UUID string `json:"uuid"`
	}{
		UUID: UUID,
	}
	schema := gojsonschema.NewStringLoader(`
	{
		"$schema": "http://json-schema.org/draft-07/schema#",
		"type": "object",
		"properties": {
			"uuid": { 
				"type": "string", 
				"format": "uuid"
			}
		}
	}`)
	loader := gojsonschema.NewGoLoader(uuid)

	return validate(schema, loader)
}

// ValidateAccount validates account.
func ValidateAccount(accountJSON []byte) (bool, error) {
	schema := gojsonschema.NewStringLoader(AccountSchema)
	loader := gojsonschema.NewBytesLoader(accountJSON)

	return validate(schema, loader)
}

func validate(schema, loader gojsonschema.JSONLoader) (bool, error) {
	result, err := gojsonschema.Validate(schema, loader)
	if err != nil {
		return false, err
	}
	if !result.Valid() {
		errorMessage := "\n"
		for _, e := range result.Errors() {
			errorMessage = errorMessage + " " + e.String() + "\n"
		}
		return false, fmt.Errorf("Validation - %v", errorMessage)
	}

	return true, nil
}
