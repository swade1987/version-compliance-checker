package devicevalidation

import (
	"errors"
	"fmt"
)

// ErrInvalidArgument is returned when one or more arguments are invalid.
var ErrInvalidArgument = errors.New("invalid argument")

// ErrRequiredVersionNotParseable is returned when a required version is not parseable.
type ErrRequiredVersionNotParseable struct {
	Version string
}

func (e *ErrRequiredVersionNotParseable) Error() string {
	return fmt.Sprintf("The required version provided (%s) is not a valid semantic version.", e.Version)
}

// ErrCurrentVersionNotParseable is returned when a current version is not parseable.
type ErrCurrentVersionNotParseable struct {
	Version string
}

func (e *ErrCurrentVersionNotParseable) Error() string {
	return fmt.Sprintf("The current version provided (%s) is not a valid semantic version.", e.Version)
}
