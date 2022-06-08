package math

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVar(t *testing.T) {
	is := assert.New(t)
	xs := []float64{5, 6, 3}
	r, err := Var(xs)
	is.Nil(err)
  is.InDelta(2.333333, r, 1e-6)
}
