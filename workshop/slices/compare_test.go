package slices

import (
	"fmt"
	"testing"
)

func TestSliceCompare(t *testing.T) {
	numbers1 := []int{1, 5, 2, 6, 7, 3, 8, 9, 3}
	numbers2 := []int{1, 5, 2, 6, 7, 3, 8, 9}
	numbers3 := []int{1, 5, 2, 6, 7, 3, 8, 9, 1}

	if !Compare(numbers1, numbers1) {
		t.Fatal(fmt.Sprintf("Expected %v to match %v", numbers1, numbers1))
	}

	if Compare(numbers1, numbers2) {
		t.Fatal(fmt.Sprintf("Expected %v got %v. Different content.", numbers1, numbers2))
	}

	if Compare(numbers1, numbers3) {
		t.Fatal(fmt.Sprintf("Expected %v got %v. Different lengths.", numbers1, numbers3))
	}
}
