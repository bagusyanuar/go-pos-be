package exception

import "errors"

var (
	ErrUserNotFound   = errors.New("user not found")
	ErrRecordNotFound = errors.New("data not found")
)
