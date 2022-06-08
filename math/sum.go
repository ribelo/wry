package math 

import "github.com/ribelo/wry"

// Returns the sum of an array.
func Sum[T wry.Number](xs []T) T {

	sum := T(0.0)

	for i := 0; i < len(xs); i++ {
		sum += xs[i]
	}

	return sum
}
