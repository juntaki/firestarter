package domain

import (
	"gopkg.in/go-playground/validator.v9"
)

type Validator struct {
	validation *validator.Validate
}

func NewValidator() *Validator {
	validate := validator.New()
	validate.RegisterStructValidation(ConfigValidator, Config{})
	return &Validator{
		validation: validate,
	}
}

func (v *Validator) ValidateConfig(config *Config) error {
	return v.validation.Struct(config)
}
