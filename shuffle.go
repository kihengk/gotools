package tools

import (
	"math/rand"
	"time"
)

func Shuffle[T any](a []T) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(a), func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})
}
