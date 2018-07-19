package functions

// Mutate ...
func Mutate(s []int) {
	i := 0
	for _, x := range s {
		if x%3 != 0 && x%5 != 0 {
			s[i] = x
			i++
		}
	}
	s = s[:i]
}
