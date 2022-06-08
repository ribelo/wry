package math

import (
	"errors"
	"fmt"
	"github.com/ribelo/wry"
	"math"
)

// Returns the correlation coefficient between two arrays.
func Cor[T wry.Number](xs []T, ys []T) (float64, error) {
	if len(xs) != len(ys) {
		return math.NaN(), errors.New(fmt.Sprintf("Arrays have different length: x[%d], y[%d]", len(xs), len(ys)))
	}

	if len(xs) < 3 {
		return math.NaN(), errors.New(fmt.Sprintf("Array length has to be at last 3."))
	}

	sxy, err := Cov(xs, ys)
	if err != nil {
		return math.NaN(), err
	}

	sxx, err := Var(xs)
	if err != nil {
		return math.NaN(), err
	}

	syy, err := Var(ys)
	if err != nil {
		return math.NaN(), err
	}

	if (sxx == 0) || (syy == 0) {
		return math.NaN(), errors.New(fmt.Sprintf("Variance is zero."))
	}

	return sxy / math.Sqrt(sxx*syy), nil

}
