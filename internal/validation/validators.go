package validation

import (
	"fmt"

	"github.com/xeipuuv/gojsonschema"
)

// ValidateURL validates URL.
func ValidateURL(url string) (bool, error) {
	urlJSON := struct {
		URL string `json:"url"`
	}{
		URL: url,
	}
	schema := gojsonschema.NewStringLoader(URLSchema)
	loader := gojsonschema.NewGoLoader(urlJSON)

	return validate(schema, loader)
}

// ValidateUUID validates uuid.
func ValidateUUID(uuid string) (bool, error) {
	uuidJSON := struct {
		UUID string `json:"uuid"`
	}{
		UUID: uuid,
	}
	schema := gojsonschema.NewStringLoader(UUIDSchema)
	loader := gojsonschema.NewGoLoader(uuidJSON)

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

		return false, fmt.Errorf("%v", errorMessage)
	}

	return true, nil
}
