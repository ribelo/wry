package math

import (
  "github.com/stretchr/testify/assert"
  "testing"
  )

func TestEntropy(t *testing.T) {
  is := assert.New(t)
  xs := []float64{-2.1968219, -0.9559913, -0.0431738, 1.0567679, 0.3853515}
  r, err := Entropy(xs)
  is.Nil(err)
  print("sex", r)
}
