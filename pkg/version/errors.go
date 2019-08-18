package version

import "fmt"

type RequiredVersionNotParseableError struct {
	Version string
}
type CurrentVersionNotParseableError struct {
	Version string
}

func (e *RequiredVersionNotParseableError) Error() string {
	return fmt.Sprintf("The required version provided (%s) is not a valid semantic version.", e.Version)
}

func (e *CurrentVersionNotParseableError) Error() string {
	return fmt.Sprintf("The current version provided (%s) is not a valid semantic version.", e.Version)
}
