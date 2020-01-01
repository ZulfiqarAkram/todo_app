package validator

import (
	"gopkg.in/go-playground/validator.v9"
)

type Validator interface {
	ValidateStruct(d interface{}) error
	RegisterCustomValidations()
}
type vStruct struct {
	v *validator.Validate
}

func (vStr *vStruct) ValidateStruct(d interface{}) error {
	return vStr.v.Struct(d)
}

func New() Validator {
	MyValidator := &vStruct{
		v: validator.New(),
	}
	MyValidator.RegisterCustomValidations()
	return MyValidator

}

//Custom validation
func (vStr *vStruct) RegisterCustomValidations() {
	_ = vStr.v.RegisterValidation("pwdMinLenSix", func(fl validator.FieldLevel) bool {
		return len(fl.Field().String()) >= 6
	})
}
