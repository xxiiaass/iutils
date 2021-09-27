package iutils

import (
	"math/rand"
	"time"
)

// 获取两个值之间的随机数
func RandInt64(min, max int64) int64 {
	if min >= max {
		return 0
	}
	rand.Seed(time.Now().UnixNano() + time.Now().Unix())
	return rand.Int63n(max-min) + min
}

// 若size <= len , 生成常规的切片
// 若size > len , 生成切片的值为常规切片中的随机值
func RandSliceIdx(size, len int) []int {
	idxs := make([]int, len)
	for i := range idxs {
		idxs[i] = i
	}
	if size <= len {
		return idxs
	}
	ret := make([]int, size)
	for i := range ret {
		ret[i], idxs = RandSlice(idxs)
	}
	return ret
}

// 从切片中获取随机值，和剩余的其他值
func RandSlice(s []int) (int, []int) {
	idx := rand.Int() % len(s)
	ret := s[idx]
	return ret, append(s[:idx], s[idx+1:]...)
}
