package fundamentals

import (
	"testing"
)

type testCase struct {
	x        string
	y        string
	expected bool
}

func TestIsSameColour(t *testing.T) {
	testCases := []testCase{
		testCase{
			x:        "a1",
			y:        "c3",
			expected: true,
		},
		testCase{
			x:        "a1",
			y:        "h3",
			expected: false,
		},
		testCase{
			x:        "d4",
			y:        "g7",
			expected: true,
		},
		testCase{
			x:        "h3",
			y:        "h7",
			expected: true,
		},
		testCase{
			x:        "a5",
			y:        "g8",
			expected: false,
		},
	}

	for _, testCase := range testCases {
		result := isSameColour(testCase.x, testCase.y)
		if result != testCase.expected {
			t.Errorf("Expected %v got %v", testCase.expected, result)
		}
	}
}
