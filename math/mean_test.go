package math

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMean(t *testing.T) {
  is := assert.New(t)
  
  xs := []float64{1, 2, 3, 4, 5}
  r1, err1 := Mean(xs)
  is.Equal(nil, err1)
  is.Equal(3.0, r1)

  ys := []float64{}
  r2, err2 := Mean(ys)
  is.NotNil(err2)
  is.True(math.IsNaN(r2))
}
