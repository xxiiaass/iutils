package iutils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsCrossDay(t *testing.T) {
	if !IsCrossDay(1585407600, 1585414800) {
		t.Errorf("跨天判断错误")
	}
	if IsCrossDay(1585328400, 1585407600) {
		t.Errorf("跨天判断错误")
	}
}

func TestSimpleName(t *testing.T) {
	if SimpleName("黑龙江省") != "黑龙江" {
		t.Errorf("省转换错误")
	}
	if SimpleName("黑龙江") != "黑龙江" {
		t.Errorf("省转换错误")
	}
	if SimpleName("厦门市") != "厦门" {
		t.Errorf("市转换错误")
	}
	if SimpleName("厦门") != "厦门" {
		t.Errorf("市转换错误")
	}
}

func TestSubStrByShowLen(t *testing.T) {
	tests := []struct {
		in     []string
		out    string
		subLen int
	}{
		{[]string{"123456789", "..."}, "123...", 6},
		{[]string{"回", "..."}, "", 1},
		{[]string{"嗷a嗷a嗷abbbb", "..."}, "嗷a...", 7},
		{[]string{"嗷a嗷a嗷abbbb", "..."}, "嗷a嗷...", 8},
	}

	for _, test := range tests {
		subStr := SubStrByShowLen(test.in[0], test.in[1], test.subLen)
		assert.Equal(t, test.out, subStr)
	}
}

func TestCeil(t *testing.T) {
	tests := []struct {
		in  []int
		out int
	}{
		{[]int{5, 4}, 8},
		{[]int{101, 100}, 200},
		{[]int{100, 100}, 100},
	}

	for _, test := range tests {
		out := Ceil(test.in[0], test.in[1])
		assert.Equal(t, test.out, out)
	}
}

func TestVersionThean(t *testing.T) {
	assert.Equal(t, VersionThan("2.6.0", "2.6.0.0"), 0)
	assert.Equal(t, VersionThan("2.6.0", "2.6.0.1"), -1)
	assert.Equal(t, VersionThan("2.6.0", "2.6.1"), -1)
	assert.Equal(t, VersionThan("2.6.1", "2.6.0.0"), 1)
	assert.Equal(t, VersionThan("2.6.0.1", "2.6.0.0"), 1)
}

func TestGetFileExt(t *testing.T) {
	assert.Equal(t, GetFileExt("daodao/daodao.mp4"), "mp4")
	assert.Equal(t, GetFileExt("daodao/daodao.."), "")
}

func TestSubstr(t *testing.T) {
	assert.Equal(t, ShowSubstr("sf4dn你d好33ma", 7), "sf4dn你")
}
