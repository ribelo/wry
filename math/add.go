package math

import (
	"errors"
  "fmt"
  "github.com/ribelo/wry"
)

func Add[T wry.Number](xs []T, ys []T) ([]T, error) {
	if len(xs) != len(ys) {
		return nil, errors.New(fmt.Sprintf("Arrays have different length: x[%d], y[%d]", len(xs), len(ys)))
	}
	for i := 0; i < len(xs); i++ {
		ys[i] = ys[i] + xs[i]
	}

	return ys, nil
}
