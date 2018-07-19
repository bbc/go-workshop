package fundamentals

import "testing"

func TestSwap(t *testing.T) {
	x, y := 100, 200
	Swap(&x, &y)

	if x != 200 || y != 100 {
		t.Errorf("Expected x = %d and y = %d. Got x = %d and y = %d", 200, 100, x, y)
	}
}
