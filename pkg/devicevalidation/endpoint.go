package devicevalidation

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type validateRequest struct {
	DeviceType     string `json:"device_type"`
	CurrentVersion string `json:"current_version"`
}

type validateResponse struct {
	Compliant       bool   `json:"compliant"`
	RequiredVersion string `json:"required_version,omitempty"`
	Err             error  `json:"error,omitempty"`
}

func (r validateResponse) error() error { return r.Err }

func makeValidateEndpoint(dvs Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(validateRequest)
		compliant, requiredVersion, err := dvs.Validate(req.DeviceType, req.CurrentVersion)
		return validateResponse{
			Compliant:       compliant,
			RequiredVersion: requiredVersion,
			Err:             err,
		}, nil
	}
}
