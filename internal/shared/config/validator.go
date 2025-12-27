package config

import (
	"github.com/bagusyanuar/go-pos-be/pkg/util"
	"github.com/go-playground/validator/v10"
)

func NewValidator() *validator.Validate {
	v := validator.New()

	util.RegisterValidatorTag(v)
	util.RegisterValidatorRule(v)
	util.RegisterValidatorTranslation(v)
	return v
}
