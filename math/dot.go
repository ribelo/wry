package math

import (
	"errors"
	"fmt"
	"github.com/ribelo/wry"
	"math"
)

// Return the dot product of two arrays.
func Dot[T wry.Number](xs []T, ys []T) (float64, error) {
	if len(xs) != len(ys) {
		return math.NaN(), errors.New(fmt.Sprintf("Arrays have different length: x[%d], y[%d]", len(xs), len(ys)))
	}

	sum := 0.0
	for i := 0; i < len(xs); i++ {
		sum += float64(xs[i]) * float64(ys[i])
	}

	return sum, nil
}
