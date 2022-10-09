package tools

func Filter[T any](a []T, filterFunc func(val T) bool) []T {
	var result []T

	for _, e := range a {
		if filterFunc(e) {
			result = append(result, e)
		}
	}
	return result
}
