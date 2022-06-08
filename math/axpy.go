package math

import (
  "errors"
	"fmt"
  "github.com/ribelo/wry"
  )

// Update an array by adding a multiple of another array ys =  * xs + ys.
func Axpy[T wry.Number](a T, xs []T, ys *[]T) (*[]T, error) {
	if len(xs) != len(*ys) {
		return nil, errors.New(fmt.Sprintf("Arrays have different length: x[%d], y[%d]", len(xs), len(*ys)))
	}

	for i := 0; i < len(xs); i++ {
		(*ys)[i] += a * xs[i]
	}

	return ys, nil
}
