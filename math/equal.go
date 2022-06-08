package math

import (
	"github.com/ribelo/wry"
  "math"
)

// Check if xs elements are equal to ys elements.
func Equal[T wry.Number](xs []T, ys []T) (bool, error) {
	if len(xs) != len(ys) {
		return false, nil
	}

	absxs := Abs(xs)
	absys := Abs(ys)
	sumabsxs := Sum(absxs)
	sumabsys := Sum(absys)

	return math.Abs(float64(sumabsxs-sumabsys)) < 1e-10, nil
}

