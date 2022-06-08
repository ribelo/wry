package math 

import (
  "errors"
  "github.com/ribelo/wry"
  "math"
  "fmt"
)

// Returns the covariance between two arrays.
func Cov[T wry.Number](xs []T, ys []T) (float64, error) {
	if len(xs) != len(ys) {
		return math.NaN(), errors.New(fmt.Sprintf("Arrays have different length: x[%d], y[%d]", len(xs), len(ys)))
	}

	if len(xs) < 3 {
		return math.NaN(), errors.New(fmt.Sprintf("Array length has to be at least 3."))
	}

	mx, err := Mean(xs)
	if err != nil {
		return math.NaN(), err
	}

	my, err := Mean(ys)
	if err != nil {
		return math.NaN(), err
	}

	sxy := 0.0
	for i := 0; i < len(xs); i++ {
		dx := float64(xs[i]) - mx
		dy := float64(ys[i]) - my
		sxy += dx * dy
	}

	return sxy / float64(len(xs) - 1), nil
}
