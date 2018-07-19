package pipeline

func create(names ...string) <-chan string {
	out := make(chan string)
	go func() {
		for _, name := range names {
			out <- name
		}
		close(out)
	}()
	return out
}

// IMPLEMENT
func clean()      {}
func capitalise() {}

// Uncomment

// func Pipeline(names ...string) []string {
// 	cleaned := []string{}
// 	for name := range capitalise(clean(create(names...))) {
// 		cleaned = append(cleaned, name)
// 	}
// 	return cleaned
// }
