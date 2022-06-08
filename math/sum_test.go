package math 

import (
  "github.com/stretchr/testify/assert"
  "testing"
  )

func TestSum(t *testing.T) {
  is := assert.New(t)
  xs := []int{1, 2, 3, 4, 5}
  ys := []int{}
  is.Equal(15, Sum(xs))
  is.Equal(0, Sum(ys))
}
