package devicemapping

import "errors"

// ErrInvalidArgument is returned when one or more arguments are invalid.
var ErrInvalidArgument = errors.New("invalid argument")

// ErrNotFound is returned when a devicemapping is not found.
var ErrNotFound = errors.New("device mapping not found")
