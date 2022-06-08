package math 

import (
  "testing"
  "github.com/stretchr/testify/assert"
  )

func TestAll(t *testing.T) {
  is := assert.New(t)
  xs := []bool{true, true, true, true, true}
  ys := []bool{true, true, true, true, false}
  is.Equal(true, All(xs))
  is.Equal(false, All(ys))
}
