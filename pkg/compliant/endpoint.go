package compliant

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
)

// Endpoints are exposed
type Endpoints struct {
	StatusEndpoint   endpoint.Endpoint
	ValidateEndpoint endpoint.Endpoint
}

// MakeStatusEndpoint returns the response from our service "status"
func MakeStatusEndpoint(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(statusRequest) // we really just need the request, we don't use any value from it
		s, err := srv.Status(ctx)
		if err != nil {
			return statusResponse{s}, err
		}
		return statusResponse{s}, nil
	}
}

// MakeValidateEndpoint returns the response from our service "validate"
func MakeValidateEndpoint(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(validateRequest)
		b, err := srv.Validate(ctx, req.DeviceType, req.CurrentVersion)
		if err != nil {
			return validateResponse{b, "1.1.1", err.Error()}, nil
		}
		return validateResponse{b, "0.0.0", ""}, nil
	}
}

// Status endpoint mapping
func (e Endpoints) Status(ctx context.Context) (string, error) {
	req := statusRequest{}
	resp, err := e.StatusEndpoint(ctx, req)
	if err != nil {
		return "", err
	}
	statusResp := resp.(statusResponse)
	return statusResp.Status, nil
}

// Validate endpoint mapping
func (e Endpoints) Validate(ctx context.Context, device string, version string) (bool, error) {
	req := validateRequest{DeviceType: device, CurrentVersion: version}
	resp, err := e.ValidateEndpoint(ctx, req)
	if err != nil {
		return false, err
	}
	validateResp := resp.(validateResponse)
	if validateResp.Err != "" {
		return false, errors.New(validateResp.Err)
	}
	return validateResp.Valid, nil
}
