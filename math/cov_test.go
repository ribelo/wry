package math

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCov(t *testing.T) {
	is := assert.New(t)
	xs := []float64{5, 6, 3}
	ys := []float64{0.5, -3.0, 2.3}
	r, err := Cov(xs, ys)
	is.Nil(err)
	is.InDelta(-3.833333, r, 1e-6)
}
