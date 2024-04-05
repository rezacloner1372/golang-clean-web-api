package validation

import (
	"github.com/go-playground/validator/v10"
	"github.com/rezacloner1372/golang-clean-web-api/common"
)

// ^09(1[0-9] |2[0-2] |3[0-9] |9[0-9]) [0-9]{7}$
func IranianMobileNumberValidator(fl validator.FieldLevel) bool {
	// get the value from the field we are validating
	value, ok := fl.Field().Interface().(string)
	if !ok {
		fl.Param()
		return false
	}
	return common.IranianMobileNumberValidate(value)
}
