package iutils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAsStringSlice(t *testing.T) {
	assert.Equal(t, "dsdfs", AsStringSlice(JsonDecodeAny("[\"dsdfs\"]"))[0])
	assert.Equal(t, "[]", JsonEncode(AsStringSlice(JsonDecodeAny("[dfs\"]"))))
}


func TestAsInt64Slice(t *testing.T) {
	strs := []string{
		"4345", "dsf", "xxx", "4353",
	}
	ints := AsInt64Slice(strs)
	assert.Equal(t, ints[0], int64(4345))
	assert.Equal(t, ints[1], int64(0))
	assert.Equal(t, ints[2], int64(0))
	assert.Equal(t, ints[3], int64(4353))
}
