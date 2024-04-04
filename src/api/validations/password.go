package validation

import "github.com/spf13/afero/internal/common"

func PasswordValidator(fl validator.FieldLevel) bool {
	value, ok := fl.Field().Interface().(string)
	if !ok {
		fl.Param
		return false
	}
	return common