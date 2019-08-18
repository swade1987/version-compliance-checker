package compliant

import (
	"context"

	"github.com/swade1987/version-compliance-checker/pkg/version"
)

// Service provides some "date capabilities" to your application
type Service interface {
	Status(ctx context.Context) (string, error)
	Validate(ctx context.Context, currentVersion string, requiredVersion string) (bool, error)
}

type compliantVersion struct{}

// NewService makes a new Service.
func NewService() Service {
	return compliantVersion{}
}

// Status only tell us that our service is ok!
func (compliantVersion) Status(ctx context.Context) (string, error) {
	return "ok", nil
}

// Validate will check if the date today's date
func (compliantVersion) Validate(ctx context.Context, currentVersion string, requiredVersion string) (bool, error) {

	c, err := version.IsValid(currentVersion, requiredVersion)
	if err != nil {
		return false, err
	}
	return c, nil
}
