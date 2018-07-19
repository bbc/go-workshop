package maps

import "testing"

func TestIterateMap(t *testing.T) {
	alphabet := []string{
		"alpha",
		"beta",
		"gamma",
		"delta",
		"epsilon",
	}
	expected := 0
	for _, name := range alphabet {
		expected += len(name)
	}
	sum := IterateMap()
	if sum != expected {
		t.Errorf("Expected %d got %d", expected, sum)
	}
}
