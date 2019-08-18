package version

import (
	"fmt"
	"github.com/Masterminds/semver"
)

func Valid(current, required string) (bool, error) {

	c, err := semver.NewConstraint(fmt.Sprintf(">= %s", required))

	if err != nil {
		return false, &RequiredVersionNotParseableError{Version: required}
	}

	v, err := semver.NewVersion(current)
	if err != nil {
		return false, &CurrentVersionNotParseableError{Version: current}
	}

	// Check if the version meets the constraints. The a variable will be true.
	return c.Check(v), nil
}
