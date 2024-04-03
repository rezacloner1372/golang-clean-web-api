package validation

import (
	"log"
	"regexp"

	"github.com/go-playground/validator/v10"
)

// ^09(1[0-9] |2[0-2] |3[0-9] |9[0-9]) [0-9]{7}$
func IranianMobileNumberValidator(fl validator.FieldLevel) bool {
	// get the value from the field we are validating
	value, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}
	// do your validation here
	res, err := regexp.MatchString(`^09(1[0-9]|2[0-2]|3[0-9]|9[0-9])[0-9]{7}$`, value)
	if err != nil {
		log.Print(err.Error())
	}
	return res
}
