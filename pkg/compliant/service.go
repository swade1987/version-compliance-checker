package compliant

import (
	"context"
)

// Service provides some "date capabilities" to your application
type Service interface {
	Status(ctx context.Context) (string, error)
	Validate(ctx context.Context, device string, currentVersion string) (bool, error)
}

type dateService struct{}

// NewService makes a new Service.
func NewService() Service {
	return dateService{}
}

// Status only tell us that our service is ok!
func (dateService) Status(ctx context.Context) (string, error) {
	return "ok", nil
}

// Validate will check if the date today's date
func (dateService) Validate(ctx context.Context, device string, currentVersion string) (bool, error) {
	return true, nil
}
