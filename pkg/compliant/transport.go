package compliant

import (
	"context"
	"encoding/json"
	"net/http"
)

type validateRequest struct {
	DeviceType     string `json:"device_type"`
	CurrentVersion string `json:"current_version"`
}

type validateResponse struct {
	Valid           bool   `json:"valid"`
	RequiredVersion string `json:"required_version"`
	Err             string `json:"err,omitempty"`
}

type statusRequest struct{}

type statusResponse struct {
	Status string `json:"status"`
}

func decodeValidateRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req validateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeStatusRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req statusRequest
	return req, nil
}

// Last but not least, we have the encoder for the response output
func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
