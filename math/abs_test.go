package math

import (
	"testing"
  "github.com/stretchr/testify/assert"
)

func TestAbs(t *testing.T) {
  is := assert.New(t)
  
  xs := []int{-1, 0, 1}
  r := Abs(xs)

  is.Equal([]int{1, 0, 1}, r)

}
