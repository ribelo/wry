package math

import (
	"github.com/ribelo/wry"
	"math"
)

// Shanon entropy of an array.
func Entropy[T wry.Number](xs []T) (float64, error) {
	h := 0.0

	for i := 0; i < len(xs); i++ {
		p := float64(xs[i]) / float64(len(xs))
		if p > 0 {
			h += p * math.Log(p)
		}
	}

	return -h, nil
}

