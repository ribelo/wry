package math

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
  is := assert.New(t)
  
  xs := []int{1, 2, 3, 4, 5}
  ys := []int{1, 2, 3, 4, 5}

  r, err := Add(xs, ys)

  is.Equal(nil, err)
  is.Equal([]int{2, 4, 6, 8, 10}, r)
}
