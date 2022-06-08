package math

import (
	"errors"
	"math"
	"github.com/ribelo/wry"
)

// Returns the mean of an array.
func Mean[T wry.Number](xs []T) (float64, error) {
  if len(xs) == 0 {
    return math.NaN(), errors.New("Empty array")
  }

	sum := Sum(xs)

	return float64(sum) / float64(len(xs)), nil
}
