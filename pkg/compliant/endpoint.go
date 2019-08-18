package compliant

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"

	"github.com/swade1987/version-compliance-checker/pkg/devicetype"
)

const (
	Android = "android"
	Ios     = "ios"
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
func MakeValidateEndpoint(srv Service, iosRequiredVersion string, androidRequiredVersion string) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(validateRequest)

		var validVersion = false
		var err error

		// Validate the device type provided.
		_, err = devicetype.Valid(req.DeviceType, Android, Ios)
		if err != nil {
			return validateResponse{Compliant: false, RequiredVersion: "", Err: err.Error()}, nil
		}

		// Validate the current version provided.
		if req.DeviceType == Android {
			validVersion, err = srv.Validate(ctx, req.CurrentVersion, androidRequiredVersion)
		}

		if req.DeviceType == Ios {
			validVersion, err = srv.Validate(ctx, req.CurrentVersion, iosRequiredVersion)
		}

		// Build the correct response.
		if err != nil {
			return validateResponse{validVersion, "", err.Error()}, nil
		}

		if validVersion == true {
			return validateResponse{validVersion, "", ""}, nil
		}

		if req.DeviceType == Android {
			return validateResponse{validVersion, androidRequiredVersion, ""}, nil
		}

		return validateResponse{validVersion, iosRequiredVersion, ""}, nil
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
	return validateResp.Compliant, nil
}
