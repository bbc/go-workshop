package functions

import (
	"reflect"
	"testing"
)

func TestMutate(t *testing.T) {
	numbers := []int{333, 999, 17, 18, 1234, 20}
	expected := []int{17, 1234}
	Mutate(numbers)

	if !reflect.DeepEqual(numbers, expected) {
		t.Fatalf("Expected %v got %v", expected, numbers)
	}
}
