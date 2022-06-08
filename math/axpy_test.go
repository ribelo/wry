package math

import (
  "github.com/stretchr/testify/assert"
  "testing"
  )

func TestAxpy(t *testing.T) {
  is := assert.New(t)
  a := 10.0
  xs := []float64{1, 2, 3, 4, 5}
  ys := []float64{1, 2, 3, 4, 5}
  r, err := Axpy(a, xs, &ys)
  is.Equal(nil, err)
  is.Equal([]float64{11, 22, 33, 44, 55}, *r)
}
