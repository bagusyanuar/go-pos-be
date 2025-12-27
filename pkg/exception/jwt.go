package exception

import "errors"

var (
	ErrTokenMissingOrMalformed = errors.New("token is missing or malformed")
	ErrTokenExpired            = errors.New("token is expired")
	ErrClaimToken              = errors.New("token cannot be claim")
	ErrInvalidSubjectFormat    = errors.New("invalid subject format")
)
