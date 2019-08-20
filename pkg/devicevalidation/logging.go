package devicevalidation

import (
	"time"

	"github.com/go-kit/kit/log"
)

type loggingService struct {
	logger log.Logger
	Service
}

// NewLoggingService returns a new instance of a logging Service.
func NewLoggingService(logger log.Logger, s Service) Service {
	return &loggingService{logger, s}
}

func (s *loggingService) Validate(deviceType string, currentVersion string) (compliant bool, requiredAppVersion string, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "validate",
			"deviceType", deviceType,
			"currentversion", currentVersion,
			"took", time.Since(begin),
			"err", err)
	}(time.Now())
	return s.Service.Validate(deviceType, currentVersion)
}
