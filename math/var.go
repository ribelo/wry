package math

import (
	"errors"
	"fmt"
	"github.com/ribelo/wry"
	"math"
)

// Returns the variane of a array.
func Var[T wry.Number](xs []T) (float64, error) {
	if len(xs) < 2 {
		return math.NaN(), errors.New(fmt.Sprintf("Array length has to be at least 2."))
	}

	sum := 0.0
	sumsq := 0.0

	for i := 0; i < len(xs); i++ {
		sum += float64(xs[i])
		sumsq += float64(xs[i] * xs[i])
	}

	return (sumsq - sum * sum/float64(len(xs))) / float64(len(xs) - 1), nil
}
