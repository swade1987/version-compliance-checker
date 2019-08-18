package devicetype

func IsValid(device, android, ios string) (bool, error) {

	allowedDeviceTypes := []string{android, ios}
	v := contains(allowedDeviceTypes, device)

	if v == false {
		return false, &NotRecognised{}
	}

	return true, nil
}

func contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}
