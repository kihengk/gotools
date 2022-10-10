package generics

import (
	"reflect"
	"testing"
)

func TestMap(t *testing.T) {
	target := []string{"a", "bc", "def"}
	result := Map(target, func(t string) int {
		return len(t)
	})

	expected := []int{1, 2, 3}
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected: %v, got: %v", expected, result)
	}
}
