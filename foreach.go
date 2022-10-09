package tools

func ForEach[T any](a []T, f func(T)) {
	for _, e := range a {
		f(e)
	}
}
