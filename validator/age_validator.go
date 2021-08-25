package validator

import "github.com/go-playground/validator/v10"

func ValidAgeValidator(minAge int64) validator.Func {
	return func(fl validator.FieldLevel) bool {
		age, ok := fl.Field().Interface().(int64)
		if ok {
			return age >= minAge
		}
		return false
	}
}
