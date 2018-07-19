package arrays

import "testing"

var testcases = map[string]rune{
	"aabbbbcddde00":  'c',
	"adbcabde0d0bb":  'c',
	"aprdfkjkksdwea": 'p',
	"zazbb":          'a',
}

func TestFindsFirstNonRepeated(t *testing.T) {
	for k, v := range testcases {
		actual := FirstNonRepeated(k)
		if actual != v {
			t.Fatalf("Expected %v, got %v", v, actual)
		}
	}
}

func TestReturnsNegativeOneWhenNoCharsRepeated(t *testing.T) {
	actual := FirstNonRepeated("aaddd00")
	if actual != -1 {
		t.Fatalf("Expected %v, got %v", -1, actual)
	}
}
