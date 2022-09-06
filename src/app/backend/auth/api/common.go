package api

func ToAuthenticateModes(modes []string) AuthenticationModes {
	result := AuthenticationModes{}
	modesMap := map[string]bool{}

	for _, mode := range []AuthenticationMode{} {
		modesMap[mode.String()] = true
	}

	return result
}
