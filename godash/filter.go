package godash

func Filter[T comparable](arr []T, selector func(a T) bool) []T {
	filtered := make([]T, 0)

	for _, v := range arr {
		if selector(v) {
			filtered = append(filtered, v)
		}
	}

	return filtered
}
