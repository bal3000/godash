package godash

func Contains[T comparable](arr []T, value T) bool {
	for _, a := range arr {
		if a == value {
			return true
		}
	}

	return false
}
