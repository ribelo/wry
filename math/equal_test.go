package math 

import (
  "github.com/stretchr/testify/assert"
  "testing"
  )

func TestEqual(t *testing.T) {
  is := assert.New(t)
  xs := []float64{-2.1968219, -0.9559913, -0.0431738, 1.0567679, 0.3853515}
  ys := []float64{-1.7781325, -0.6659839, 0.9526148, -0.9460919, -0.3925300}
  zs := []float64{-1.7781325, -0.6659839, 0.9526148, -0.9460919, -0.3925300}
  r1, err := Equal(xs, ys)
  r2, err := Equal(ys, zs)
  is.Nil(err)
  is.False(r1)
  is.True(r2)
}
