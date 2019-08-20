package devicemapping

// DeviceMapping maps a device type to a required app version.
type DeviceMapping struct {
	DeviceType         string `json:"deviceType"`
	RequiredAppVersion string `json:"requiredAppVersion"`
}

// Repository stores device mappings.
type Repository interface {
	Find(deviceType string) (*DeviceMapping, error)
	AddRequiredVersion(deviceType string, requiredAppVersion string) error
}

// inMemRepository stores device mappings in memory.
type inMemRepository struct {
	DeviceMappings map[string]*DeviceMapping `json:"deviceMappings"`
}

// NewInMemRepository creates a new in memory repository.
func NewInMemRepository() Repository {
	return &inMemRepository{
		DeviceMappings: map[string]*DeviceMapping{},
	}
}

// Find a device mapping by device type.
func (r *inMemRepository) Find(deviceType string) (*DeviceMapping, error) {
	if dm, ok := r.DeviceMappings[deviceType]; ok {
		return dm, nil
	} else {
		return nil, ErrNotFound
	}
}

// AddRequiredVersion a device type mapping
func (r *inMemRepository) AddRequiredVersion(deviceType string, requiredAppVersion string) error {
	if len(deviceType) == 0 {
		return ErrInvalidArgument
	}

	r.DeviceMappings[deviceType] = &DeviceMapping{
		DeviceType:         deviceType,
		RequiredAppVersion: requiredAppVersion,
	}

	return nil
}
