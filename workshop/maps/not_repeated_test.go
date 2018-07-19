package maps

import "testing"

func TestFindsFirstNonRepeated(t *testing.T) {
	expected := 'c'
	actual := FirstNonRepeated("aabbbbcddde00")
	if actual != expected {
		t.Fatalf("Expected %v, got %v", expected, actual)
	}
}

func TestReturnsNegativeOneWhenNoCharsRepeated(t *testing.T) {
	actual := FirstNonRepeated("aaddd00")
	if actual != -1 {
		t.Fatalf("Expected %v, got %v", -1, actual)
	}
}
