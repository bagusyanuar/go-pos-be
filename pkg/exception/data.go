package exception

import "errors"

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrRecordNotFound     = errors.New("data not found")
	ErrUnitConversionRate = errors.New("default unit conversion rate must be 1")
	ErrUnitDefault        = errors.New("unit must have exactly one default unit")
	ErrUnitDefaultValue   = errors.New("is default must have boolean value")
	ErrDeleteDefaultUnit  = errors.New("cannot delete. this unit is default")
)
