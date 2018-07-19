package testing

import "errors"

func getDigits(n int) (int, int) {
	digits := []int{}
	x := n
	for x > 0 {
		y := x % 10
		digits = append([]int{y}, digits...)
		x /= 10
	}
	return digits[0], digits[1]
}

// Magic9 ...
func Magic9(n int) (int, error) {
	if n < 20 || n > 99 {
		return -1, errors.New("Number must be between 20 and 99 (inclusive)")
	}
	x, y := getDigits(n)
	x = n - (x + y)
	x, y = getDigits(x)
	return x + y, nil
}
