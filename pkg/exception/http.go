package exception

import "errors"

var (
	ErrPasswordMissmatch = errors.New("password did not match")
	ErrRouteNotFound     = errors.New("route not found")
)
