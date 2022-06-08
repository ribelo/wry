package math

import (
	"errors"
	"fmt"
	"github.com/ribelo/wry"
	"math"
)

// The Squared Euclidean distance between two arrays.
func SqDist[T wry.Number](xs []T, ys []T) (float64, error) {
	if len(xs) != len(ys) {
		return math.NaN(), errors.New(fmt.Sprintf("Arrays have different length: x[%d], y[%d]", len(xs), len(ys)))
	}

	switch len(xs) {
	case 2:
		d0 := xs[0] - ys[0]
		d1 := xs[1] - ys[1]
		return float64(d0*d0 + d1*d1), nil
	case 3:
		d0 := xs[0] - ys[0]
		d1 := xs[1] - ys[1]
		d2 := xs[2] - ys[2]
		return float64(d0*d0 + d1*d1 + d2*d2), nil
	case 4:
		d0 := xs[0] - ys[0]
		d1 := xs[1] - ys[1]
		d2 := xs[2] - ys[2]
		d3 := xs[3] - ys[3]
		return float64(d0*d0 + d1*d1 + d2*d2 + d3*d3), nil
	}

	sum := 0.0
	for i := 0; i < len(xs); i++ {
		d := xs[i] - ys[i]
		sum += float64(d * d)
	}

	return sum, nil
}
