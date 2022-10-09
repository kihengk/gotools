package tools

func Map[T any, R any](a []T, f func(T) R) []R {
	var result []R
	for _, e := range a {
		val := f(e)
		result = append(result, val)
	}
	return result
}
