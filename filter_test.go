package tools

import (
	"reflect"
	"testing"
)

func TestFilter(t *testing.T) {
	target := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := []int{2, 4, 6, 8, 10}

	result := Filter(target, func(t int) bool {
		return t%2 == 0
	})

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected: %v, got: %v", expected, result)
	}
}
