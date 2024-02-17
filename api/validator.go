package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/yyht/simplebank/util"
)

// register this custom register with Gin

// validator.Func is of type func(fl FieldLevel) bool
// fl is an interface which contains all information and helper functinos to validate a field
var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {

	// fieldLevel.Field() to get reflection value of the field
	// .Interface(), to get this value as an empty interface
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		// check currency is supported
		return util.IsSupportedCurrency(currency)
	}

	return false
}
