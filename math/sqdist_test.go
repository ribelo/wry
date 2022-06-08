package math

import (
  "github.com/stretchr/testify/assert"
  "testing"
  )

func TestSqDist(t *testing.T) {
  is := assert.New(t)
  xs := []float64{-2.1968219, -0.9559913, -0.0431738, 1.0567679, 0.3853515}
  ys := []float64{-1.7781325, -0.6659839, 0.9526148, -0.9460919, -0.3925300}
  r, err := SqDist(xs, ys)
  is.Nil(err)
  is.InDelta(5.867547, r, 1e-6)
}
