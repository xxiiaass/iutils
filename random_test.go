package iutils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRandSliceIdx(t *testing.T) {
	RandSliceIdx(3, 3)
}

func TestRandSlice(t *testing.T) {
	ints := []int{13, 5, 2}
	num, slice := RandSlice([]int{13, 5, 2})
	assert.Equal(t, 2, len(slice))
	assert.Contains(t, ints, num)
}
