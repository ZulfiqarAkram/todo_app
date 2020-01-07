package validator

import (
	"gopkg.in/go-playground/validator.v9"
)

type Validator interface {
	ValidateStruct(d interface{}) error
	RegisterCustomValidations()
}
type Validate struct {
	validator *validator.Validate
}

func (v *Validate) ValidateStruct(d interface{}) error {
	return v.validator.Struct(d)
}

func New() *Validate {
	MyValidator := &Validate{
		validator: validator.New(),
	}
	MyValidator.RegisterCustomValidations()
	return MyValidator

}

//Custom validation
func (v *Validate) RegisterCustomValidations() {
	_ = v.validator.RegisterValidation("pwdMinLenSix", func(fl validator.FieldLevel) bool {
		return len(fl.Field().String()) >= 6
	})
}
