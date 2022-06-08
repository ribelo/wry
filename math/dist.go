package math

import ( 
  "github.com/ribelo/wry"
  "math"
  )

// The Euclidean distance between two arrays.
func Dist[T wry.Number](xs []T, ys []T) (float64, error) {
	sqd, err := SqDist(xs, ys)
	if err != nil {
		return math.NaN(), err
	}

	return math.Sqrt(sqd), nil
}
