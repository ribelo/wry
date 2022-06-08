package math 

import (
  "github.com/stretchr/testify/assert"
  "testing"
  )

func TestCor(t *testing.T) {
  is := assert.New(t)
  xs := []float64{5, 6, 3}
  ys := []float64{0.5, -3, 2.3}
  r, err := Cor(xs, ys)
  is.Nil(err)
  is.InDelta(-0.931151, r, 1e-6)
}


