package validator

import (
	"gopkg.in/go-playground/validator.v9"
)

type Validator interface {
	ValidateStruct(d interface{}) error
	RegisterCustomValidations()
}
type VStruct struct {
	v *validator.Validate
}

func (vStr *VStruct) ValidateStruct(d interface{}) error {
	return vStr.v.Struct(d)
}

func New() *VStruct {
	MyValidator := &VStruct{
		v: validator.New(),
	}
	MyValidator.RegisterCustomValidations()
	return MyValidator

}

//Custom validation
func (vStr *VStruct) RegisterCustomValidations() {
	_ = vStr.v.RegisterValidation("pwdMinLenSix", func(fl validator.FieldLevel) bool {
		return len(fl.Field().String()) >= 6
	})
}
