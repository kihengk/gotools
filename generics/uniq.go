package generics

func Unique[T comparable](a []T) []T {
	u := make([]T, 0, len(a))
	m := make(map[T]struct{})

	for _, v := range a {
		if _, ok := m[v]; !ok {
			m[v] = struct{}{}
			u = append(u, v)
		}
	}
	return u
}
