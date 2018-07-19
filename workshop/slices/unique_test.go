package slices

import "testing"

func TestIsUnique(t *testing.T) {
	unqiue := "abcdef"
	notUnique := "abbcdef"

	actual1 := IsUnique(unqiue)
	actual2 := IsUnique(notUnique)

	if !actual1 {
		t.Fatalf("String %s, contains unique characters", unqiue)
	}

	if actual2 {
		t.Fatalf("String %s, does not contain unique characters", notUnique)
	}
}
