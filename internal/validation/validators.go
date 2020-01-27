package validation

import (
	"github.com/asaskevich/govalidator"
)

func ValidateURL(URL string) bool {
	return govalidator.IsURL(URL)
}

func ValidateUUID(UUID string) bool {
	return govalidator.IsUUID(UUID)
}
