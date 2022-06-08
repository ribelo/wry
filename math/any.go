package math

// Given a set of boolean values, is at least one of the values true?
func Any(xs []bool) bool {
	for i := 0; i < len(xs); i++ {
		if xs[i] {
			return true
		}
	}

	return false
}
