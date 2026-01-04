package exception

import "errors"

var (
	ErrPasswordMissmatch      = errors.New("password did not match")
	ErrRouteNotFound          = errors.New("route not found")
	ErrInvalidQueryParameters = errors.New("invalid query parameters")
	ErrInvalidRequestBody     = errors.New("invalid request body")
	ErrValidation             = errors.New("validation error")
	ErrUnitConversionRate     = errors.New("default unit conversion rate must be 1")
	ErrUnitDefault            = errors.New("unit must have exactly one default unit")
)
