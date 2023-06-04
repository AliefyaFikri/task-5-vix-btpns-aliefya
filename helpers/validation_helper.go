package helpers

import (
	"fmt"
	"strings"

	"github.com/asaskevich/govalidator"
)

// ValidateStruct validates a struct using the "valid" tag
func ValidateStruct(s interface{}) error {
	if _, err := govalidator.ValidateStruct(s); err != nil {
		validationErrors := err.(govalidator.Errors)
		errorMessages := make([]string, len(validationErrors))
		for i, e := range validationErrors {
			errorMessages[i] = e.Error()
		}
		return fmt.Errorf(strings.Join(errorMessages, ", "))
	}
	return nil
}
