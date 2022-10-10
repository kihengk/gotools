package tools

import "testing"

func BenchmarkUnique(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Unique([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1})
	}
}
