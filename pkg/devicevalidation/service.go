// Package devicevalidation provides the use-case of matching a device
// to a valid version number.
package devicevalidation

import (
	"fmt"

	"github.com/Masterminds/semver"

	"github.com/swade1987/version-compliance-checker/pkg/devicemapping"
)

// Service is the interface that provides the Validate method.
type Service interface {
	// Validate a deviceType and its current version. Return the compliance
	// state and required version to be compliant.
	Validate(deviceType string, currentVersion string) (bool, string, error)
}

type service struct {
	deviceMappings devicemapping.Repository
}

func (s *service) Validate(deviceType string, currentVersion string) (bool, string, error) {
	if len(deviceType) == 0 {
		return false, "", ErrInvalidArgument
	}
	dm, err := s.deviceMappings.Find(deviceType)
	if err != nil {
		return false, "", err
	}

	constraint, err := semver.NewConstraint(fmt.Sprintf(">= %s", dm.RequiredAppVersion))
	if err != nil {
		return false, "", &ErrRequiredVersionNotParseable{Version: dm.RequiredAppVersion}
	}

	compareVersion, err := semver.NewVersion(currentVersion)
	if err != nil {
		return false, "", &ErrCurrentVersionNotParseable{Version: currentVersion}
	}

	// Check if the version meets the constraints.
	if constraint.Check(compareVersion) {
		return true, "", nil
	}

	return false, dm.RequiredAppVersion, nil
}

// NewService returns a new instance of the default Service.
func NewService(deviceMappings devicemapping.Repository) Service {
	return &service{
		deviceMappings: deviceMappings,
	}
}
