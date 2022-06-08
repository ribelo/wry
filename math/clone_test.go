package math

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClone(t *testing.T) {
  is := assert.New(t)
  xs := []int{1, 2, 3, 4, 5}
  ys := Clone(xs)
  is.Equal(xs, ys)
}

func TestClone2d(t *testing.T) {
  is := assert.New(t)
  xs := [][]int{{1, 2, 3}, {4, 5, 6}}
  ys := Clone2d(xs)
  is.Equal(xs, ys)
}
