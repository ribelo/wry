package math

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestDist(t *testing.T) {
  is := assert.New(t)
  xs := []float64{-2.1968219, -0.9559913, -0.0431738, 1.0567679, 0.3853515}
  ys := []float64{-1.7781325, -0.6659839, 0.9526148, -0.9460919, -0.3925300}
  r, err := Dist(xs, ys)
  is.Nil(err)
  is.InDelta(2.422302, r, 1e-6)
}
