package math

// Given a set of boolean values, are all of the values true?
func All(xs []bool) bool {

	for i := 0; i < len(xs); i++ {
		if !xs[i] {
			return false
		}
	}

	return true
}
