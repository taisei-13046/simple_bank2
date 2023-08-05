package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/taisei-13046/simple_bank2/util"
)

var validCurrency validator.Func = func(fiedlLevel validator.FieldLevel) bool {
	if currency, ok := fiedlLevel.Field().Interface().(string); ok {
		return util.IsSupportedCurrency(currency)
	}
	return false
}
