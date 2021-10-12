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

func TestStrToBytes(t *testing.T) {
	s := "abcd"

	bs := StrToBytes(s)
	assert.Equal(t, uint8('a'), bs[0])
	assert.Equal(t, uint8('d'), bs[3])

	bs[1] = 's'
	assert.Equal(t, uint8('s'), bs[1])

	unsafeBs := StrToBytes(s, true)

	assert.Equal(t, uint8('a'), unsafeBs[0])
	assert.Equal(t, uint8('d'), unsafeBs[3])

	// 尝试修改，会报错
	//unsafeBs[1] = 'u'
	//assert.Equal(t, uint8('u'), bs[1])
}

func TestBytesToStr(t *testing.T) {
	bs := []byte{'a', 'b', 'c', 'd'}

	s := BytesToStr(bs)
	assert.Equal(t, "abcd", s)

	unsafeS := BytesToStr(bs, true)

	assert.Equal(t, "abcd", unsafeS)

	// 此时修改原byte数组，会导致unsafeS改变
	bs[2] = 'u'
	assert.Equal(t, "abud", unsafeS)
}
