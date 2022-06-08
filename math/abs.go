package math

import (
  "math"
  "github.com/ribelo/wry"
)

// Return asbolute values of a array.
func Abs[T wry.Number](xs []T) ([]T) {
	ys := make([]T, len(xs))
	for i := 0; i < len(xs); i++ {
		ys[i] = T(math.Abs(float64(xs[i])))
	}

	return ys
}
