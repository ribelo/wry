package math

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAny(t *testing.T) {
	is := assert.New(t)
	xs := []bool{true, true, true, true, true}
	ys := []bool{true, true, true, true, false}
	zs := []bool{false, false, false, false, false}
	is.Equal(true, Any(xs))
	is.Equal(true, Any(ys))
	is.Equal(false, Any(zs))
}
